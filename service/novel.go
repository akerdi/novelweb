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
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

func SearchRecommandNovel() gin.HandlerFunc {
	return func(context *gin.Context) {
		keyWord := context.Param("keyword")
		dbFD := db.GetDB()
		likeStr := fmt.Sprintf("%%%s%%", keyWord)
		var novelNets []schema.NovelNet
		dbc := dbFD.Where("title LIKE ?", likeStr).Limit(15).Find(&novelNets)
		if dbc.Error != nil {
			fmt.Println("database meet error!!!1", dbc.Error)
			context.JSON(http.StatusServiceUnavailable, nil)
			return
		}
		context.JSON(http.StatusOK, novelNets)
	}
}

// SearchNovel 搜索关键字找到对应网站
func SearchNovel() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyWord := c.Param("keyword")
		var page = "1"
		page = c.Param("page")
		var totalNovels []schema.NovelNet
		// 搜到小说，还继续搜索网络。同时保存
		_, results := searchByNet(keyWord, page)
		dbFD := db.GetDB()
		for _, result := range results {
			newNovel := schema.NovelNet{
				IsParse: uint64(result.IsParse),
				URL:     result.Href,
				Title:   result.Title,
			}
			hashStr := fmt.Sprintf("%s%s", newNovel.URL, newNovel.Title)
			newNovel.MD5 = generate.GetMD5Hash(hashStr)
			totalNovels = append(totalNovels, newNovel)
			dbc := dbFD.Create(&newNovel)
			if dbc.Error != nil {
				log.Println("SearchNovel save newNovel err:", dbc.Error.Error())
				continue
			}
		}
		c.JSON(http.StatusOK, totalNovels)
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
			fmt.Println("[novel.SearchChapter] get novelChapter err: ", dbc.Error)
			c.JSON(http.StatusOK, gin.H{
				"chapter": novelChapter,
			})
			return
		}

		var novelNet schema.NovelNet
		dbc = dbFD.Where("md5 = ?", md5).First(&novelNet)
		if dbc.Error != nil {
			fmt.Println("[novel.SearchChapter] get novelNet err:", dbc.Error)
			c.JSON(http.StatusServiceUnavailable, dbc.Error)
			return
		}
		//if (schema.NovelNet{}) == novelNet {
		//	c.JSON(http.StatusNoContent, "失败")
		//	return
		//}
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
		c.JSON(http.StatusOK, gin.H{
			"chapter": novelChapterRes,
		})
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
			log.Println("[novel.SearchContent] err: ", err)
			context.JSON(http.StatusServiceUnavailable, err)
			return
		}
		uchapterIndex := uint64(chapterIndex)
		md5 := context.Param("md5")
		log.Println("[novel.SearchContent] index: md5: ", chapterIndex, md5)
		var novelChapter schema.NovelChapter
		dbFD := db.GetDB()
		dbc := dbFD.Where("md5 = ?", md5).First(&novelChapter)
		if dbc.Error != nil {
			log.Println("[novel.SearchContent] dbcError:", dbc.Error)
			context.JSON(http.StatusServiceUnavailable, "")
			return
		}
		//这里开始寻找有没有对应章节
		var novelContentLocal schema.NovelContent
		dbc = dbFD.Where("md5_index = ?", md5+":"+chapterIndexStr).First(&novelContentLocal)
		if dbc.Error == nil {
			log.Println("[novel.SearchContent] dbcError2:", dbc.Error)
			chapterElement := novelChapter.Chapters[uchapterIndex]
			context.JSON(http.StatusOK, gin.H{
				"content": novelContentLocal,
				"element": chapterElement,
				"name": novelChapter.Name,
				"originURL": novelChapter.OriginURL,
			})
			return
		}
		chapterElement := novelChapter.Chapters[uchapterIndex]
		//chapterElement := novelChapter.Chapters[uchapterIndex]
		novelContent, err := searchContent(&novelChapter, uchapterIndex)
		if err != nil {
			context.JSON(http.StatusServiceUnavailable, "")
			return
		}
		if (schema.NovelContent{}) == *novelContent {
			context.JSON(http.StatusNoContent, nil)
			return
		}
		dbc = dbFD.Create(&novelContent)
		if dbc.Error != nil {
			log.Println("###########", dbc.Error)
			context.JSON(http.StatusServiceUnavailable, nil)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"content": novelContent,
			"element": chapterElement,
			"name": novelChapter.Name,
			"originURL": novelChapter.OriginURL,
		})
	}
}

func ConfigRule() gin.HandlerFunc {
	return func(context *gin.Context) {
		config.ConfigRule()
		context.String(http.StatusOK, "OK")
	}
}

// helper

func searchContent(novelChapter *schema.NovelChapter, index uint64) (*schema.NovelContent, error) {
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
	c := generate.NewFetcher()
	requestURI, _ := url.ParseRequestURI(novelChapter.Domain)
	host := requestURI.Host
	contentSelector, _ := config.RuleConfig.Rules[host]["content_selector"].(string)
	if contentSelector == "" {
		return &novelContent, errors.New(fmt.Sprintf("%s 解析 %s 网址 没有找到 content_selector", novelChapter.Name, html))
	}
	c.OnHTML(contentSelector, func(element *colly.HTMLElement) {
		htmlContent, err := element.DOM.Html()
		htmlContent = strings.TrimPrefix(htmlContent, " ")
		if err != nil {
			log.Println(fmt.Sprintf("%s 解析 %s 网址 %v", novelChapter.Name, html, err))
			return
			//return &novelContent, errors.New(fmt.Sprintf("%s 解析 %s 网址 %v", novelChapter.Name, html, err))
		}
		if htmlContent != "" {
			novelContent.Content = htmlContent
			novelContent.MD5Index = fmt.Sprintf("%s:%d", novelChapter.MD5, index)
			novelContent.ContentURL = html
		}
	})
	err := c.Visit(html)
	fmt.Println("[novel.searchContent] html: ", html, "selector: ", contentSelector)
	return &novelContent, err
}

func searchChapter(novelNet *schema.NovelNet) (*schema.NovelChapter, error) {
	var novelChapter schema.NovelChapter
	c := generate.NewFetcher()
	requestURI, err := url.ParseRequestURI(novelNet.URL)
	if err != nil {
		return &novelChapter, err
	}
	host := requestURI.Host
	chapterSelector, ok := config.RuleConfig.Rules[host]["chapter_selector"].(string)
	fmt.Printf("%s use chapterSelector %s\n", novelNet.Title, chapterSelector)
	if !ok {
		return &novelChapter, errors.New(fmt.Sprintf("%s in RuleConfig.Rules chapter_selector is not ok", host))
	}
	chapterLinkPrefix, ok := config.RuleConfig.Rules[host]["link_prefix"].(string)
	if !ok {
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
		log.Println("[novel.searchChapter] visit err: ", err)
		return nil, err
	}
	if len(chapterElements) > 0 {
		novelChapter.Chapters = chapterElements
		novelChapter.MD5 = novelNet.MD5
		novelChapter.LinkPrefix = chapterLinkPrefix
		novelChapter.OriginURL = novelNet.URL
		novelChapter.Name = novelNet.Title
		novelChapter.Domain = fmt.Sprintf("%s://%s", requestURI.Scheme, requestURI.Host)
	}
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
