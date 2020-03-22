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
	"sync"

	"github.com/gin-gonic/gin"
)

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

func SearchChapter() gin.HandlerFunc {
	return func(c *gin.Context) {
		md5 := c.Param("md5")
		var novelNet schema.NovelNet
		dbFD := db.GetDB()
		dbc := dbFD.Where("md5 = ?", md5).First(&novelNet)
		if dbc.Error != nil {
			fmt.Println("@@@@@@@@2", dbc.Error)
			c.JSON(http.StatusServiceUnavailable, dbc.Error)
			return
		}
		//if (schema.NovelNet{}) == novelNet {
		//	log.Println("@@@@@@@#####$$$$$$", novelNet)
		//	c.JSON(http.StatusNoContent, "失败")
		//	return
		//}
		fmt.Println("2222222", novelNet)
		novelChapter, err := searchChapter(&novelNet)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		dbc = dbFD.Create(novelChapter)
		if dbc.Error != nil {
			log.Println("-------", dbc.Error)
			c.JSON(http.StatusServiceUnavailable, "")
			return
		}
		
		//fmt.Printf("2222222333333%v\n", novelChapter)
		c.JSON(http.StatusOK, novelChapter)
	}
}

// helper

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
