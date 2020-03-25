package service

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"net/url"
	"novelweb/config"
	"novelweb/db"
	"novelweb/db/schema"
	"novelweb/generate"
	"novelweb/model"
	"novelweb/service/searchEngine"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// SearchNovel 搜索关键字找到对应网站
func SearchNovel() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyWord := c.Param("keyword")
		var page = "1"
		page = c.Param("page")
		dbFD := db.GetDB()
		var novelNets []schema.NovelNet
		likeStr := fmt.Sprintf("%%%s%%", keyWord)
		dbc := dbFD.Where("title LIKE ?", likeStr).Limit(15).Find(&novelNets)
		if dbc.Error != nil {
			fmt.Println("database meet error!!!1", dbc.Error)
			c.JSON(http.StatusServiceUnavailable, nil)
			return
		}
		if len(novelNets) > 0 {
			fmt.Println("database got data : ", novelNets)
			c.JSON(http.StatusOK, novelNets)
			return
		}
		_, results := searchByNet(keyWord, page)
		for _, result := range results {
			newNovel := schema.NovelNet{
				IsParse: uint64(result.IsParse),
				URL:     result.Href,
				Title:   result.Title,
			}
			hashStr := fmt.Sprintf("%s%s", newNovel.URL, newNovel.Title)
			newNovel.MD5 = generate.GetMD5Hash(hashStr)
			dbc = dbFD.Create(&newNovel)
			novelNets = append(novelNets, newNovel)
			if dbc.Error != nil {
				log.Println("~~~~~~~~~~~`", dbc.Error.Error())
				continue
			}
		}
		fmt.Printf("results:: %s\n", novelNets)
		c.JSON(http.StatusOK, novelNets)
	}
}

// SearchChapter 搜索章节
func SearchChapter() gin.HandlerFunc {
	return func(c *gin.Context) {
		md5 := c.Param("md5")

		dbFD := db.GetDB()
		var novelChapter schema.NovelChapter
		dbc := dbFD.Where("md5 = ?", md5).First(&novelChapter)
		if dbc.Error == nil {
			fmt.Println("~~~~~~~~~~~~~", dbc.Error)
			c.JSON(http.StatusOK, novelChapter)
			return
		}

		var novelNet schema.NovelNet
		dbc = dbFD.Where("md5 = ?", md5).First(&novelNet)
		if dbc.Error != nil {
			fmt.Println("@@@@@@@@2", dbc.Error)
			c.JSON(http.StatusServiceUnavailable, dbc.Error)
			return
		}
		//if (schema.NovelNet{}) == novelNet {
		//	c.JSON(http.StatusNoContent, "失败")
		//	return
		//}
		fmt.Println("2222222", novelNet)
		novelChapterRes, err := searchChapter(&novelNet)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		dbc = dbFD.Create(novelChapterRes)
		if dbc.Error != nil {
			c.JSON(http.StatusServiceUnavailable, "")
			return
		}
		c.JSON(http.StatusOK, novelChapterRes)
	}
}

func SearchContent() gin.HandlerFunc {
	return func(context *gin.Context) {
		chapterIndexStr := context.Param("chapterIndex")
		if chapterIndexStr == "" {
			chapterIndexStr = "1"
		}
		chapterIndex, err := strconv.Atoi(chapterIndexStr)
		if err != nil {
			log.Println("~~~~~~~", err)
			context.JSON(http.StatusServiceUnavailable, err)
			return
		}
		uchapterIndex := uint64(chapterIndex)
		md5 := context.Param("md5")
		log.Println("####3", chapterIndex, md5)
		var novelChapter schema.NovelChapter
		dbFD := db.GetDB()
		dbc := dbFD.Where("md5 = ?", md5).First(&novelChapter)
		if dbc.Error != nil {
			log.Println("@@@@@", dbc.Error)
			context.JSON(http.StatusServiceUnavailable, "")
			return
		}
		//这里开始寻找有没有对应章节
		var novelContentLocal schema.NovelContent
		dbc = dbFD.Where("md5_index = ?", md5+":"+chapterIndexStr).First(&novelContentLocal)
		if dbc.Error == nil {
			log.Println("@@@@@@@@#############!@#$%^%^^&^%*&^(*&^(&*(*&$%^%$#%$##")
			context.JSON(http.StatusOK, novelContentLocal)
			return
		}
		//chapterElement := novelChapter.Chapters[uchapterIndex]
		novelContent, err := searchContent(&novelChapter, uchapterIndex)
		if err != nil {
			context.JSON(http.StatusServiceUnavailable, "")
			return
		}
		if (schema.NovelContent{}) == *novelContent {
			context.JSON(http.StatusNoContent, "")
			return
		}
		dbc = dbFD.Create(&novelContent)
		if dbc.Error != nil {
			log.Println("###########",dbc.Error)
			context.JSON(http.StatusUnauthorized, "")
			return
		}
		context.JSON(http.StatusOK, novelContent)
	}
}

// helper

func searchContent(novelChapter *schema.NovelChapter, index uint64) (*schema.NovelContent, error)  {
	var novelContent schema.NovelContent
	novelChapterElement := novelChapter.Chapters[index]
	var html string
	if novelChapter.LinkPrefix == "1" {
		html = novelChapterElement.Href
	} else if novelChapter.LinkPrefix == "-1" {
		html = generate.UrlJoin(novelChapterElement.Href, novelChapter.OriginURL)
	} else if novelChapter.LinkPrefix == "0" {
		html = generate.UrlJoin(novelChapterElement.Href, novelChapter.Domain)
	}
	log.Println("111111", novelChapterElement)
	c:= generate.NewFetcher()
	requestURI, _ := url.ParseRequestURI(novelChapter.Domain)
	host := requestURI.Host
	contentSelector, _ := config.RuleConfig.Rules[host]["content_selector"].(string)
	if contentSelector == "" {
		return &novelContent, errors.New(fmt.Sprintf("%s 解析 %s 网址 没有找到 content_selector", novelChapter.Name, html))
	}
	c.OnHTML(contentSelector, func(element *colly.HTMLElement) {
		htmlContent, err := element.DOM.Html()
		if err != nil {
			log.Println(fmt.Sprintf("%s 解析 %s 网址 %v", novelChapter.Name, html, err))
			return
			//return &novelContent, errors.New(fmt.Sprintf("%s 解析 %s 网址 %v", novelChapter.Name, html, err))
		}
		novelContent.Content = htmlContent
		novelContent.MD5Index = fmt.Sprintf("%s:%d", novelChapter.MD5, index)
		novelContent.ContentURL = html
	})
	fmt.Println("[SearchContent] html", html)
	err := c.Visit(html)
	return &novelContent, err
}

func searchChapter(novelNet *schema.NovelNet) (*schema.NovelChapter, error) {
	var novelChapter schema.NovelChapter
	c := generate.NewFetcher()
	requestURI, err := url.ParseRequestURI(novelNet.URL)
	if err != nil {
		return &novelChapter, err
	}
	host  := requestURI.Host
	chapterSelector, ok := config.RuleConfig.Rules[host]["chapter_selector"].(string)
	if !ok {
		log.Println("11111")
		return &novelChapter, errors.New(fmt.Sprintf("%s in RuleConfig.Rules chapter_selector is not ok", host))
	}
	chapterLinkPrefix, ok := config.RuleConfig.Rules[host]["link_prefix"].(string)
	if !ok {
		log.Println("22222")
		return &novelChapter, errors.New(fmt.Sprintf("%s in RuleConfig.Rules link_prefix is not ok", host))
	}
	var chapterElements []schema.NovelChapterElement
	c.OnHTML(chapterSelector, func(element *colly.HTMLElement) {
		html := element.Attr("href")
		if html == "" {
			fmt.Println("无效dom")
			return
		}
		chapterElement := schema.NovelChapterElement{
			Name: element.Text,
			Href: html,
		}
		chapterElements = append(chapterElements, chapterElement)
	})
	fmt.Println("[chapter] html", novelNet.URL)
	err = c.Visit(novelNet.URL)
	if err != nil {
		log.Println("~~~~~~~~~~", err)
		return nil, err
	}

	novelChapter.Chapters = chapterElements
	novelChapter.MD5 = novelNet.MD5
	novelChapter.LinkPrefix = chapterLinkPrefix
	novelChapter.OriginURL = novelNet.URL
	novelChapter.Domain = fmt.Sprintf("%s://%s", requestURI.Scheme, requestURI.Host)
	log.Println("3333333", novelChapter)
	return &novelChapter, nil
}

func searchByNet(keyWord, page string) (error, []model.SearchResult) {
	// 使用网络的
	group := sync.WaitGroup{}
	//results := make([]*model.SearchResult, 0)
	results := make([]model.SearchResult, 0)
	group.Add(1)
	searchEngine := searchEngine.NewBaiduSearchEngine(func(result *model.SearchResult) {
		results = append(results, *result)
	})
	go searchEngine.EngineRun(keyWord, page, &group)
	group.Wait()
	return nil, results
}

func returnJSONModelArray() {
	var searchs = make([]model.SearchResult, 0)
	var aSearch = model.SearchResult{
		Href:    "https://www.qianpian.net/25_25871/dir.html",
		Title:   " 一代大侠全文阅读_一代大侠最新章节_snowxefd作品_千篇小说",
		IsParse: 1,
		Host:    "www.qianpian.net",
	}
	searchs = append(searchs, aSearch)
	fmt.Println("aSearch: ", aSearch)
}
