package service

import (
	"fmt"
	"log"
	"net/http"
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
		var page string = "1"
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
				IsParse: 0,
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
		c.String(http.StatusOK, "hahah"+c.Param("md5"))
	}
}

// helper

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
