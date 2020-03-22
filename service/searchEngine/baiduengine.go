package searchEngine

import (
	"fmt"
	"log"
	"net/url"
	"novelweb/config"
	"novelweb/generate"
	"novelweb/model"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type BaiduSearchEngine struct {
	parseRule       string
	searchRule      string
	domain          string
	parseResultFunc func(searResult *model.SearchResult)
}

func NewBaiduSearchEngine(parseResultFunc func(result *model.SearchResult)) *BaiduSearchEngine {
	return &BaiduSearchEngine{
		parseRule:       "#content_left h3.t a",
		searchRule:      "intitle: %s 小说 阅读",
		domain:          "http://www.baidu.com/s?wd=%s&ie=utf-8&rn=15&vf_bl=1",
		parseResultFunc: parseResultFunc,
	}
}

func (engine *BaiduSearchEngine) EngineRun(novelName, page string, group *sync.WaitGroup) {
	defer group.Done()
	searchKey := url.QueryEscape(fmt.Sprintf(engine.searchRule, novelName))
	requestUrl := fmt.Sprintf(engine.domain, searchKey)
	pageIndex, err := strconv.Atoi(page)
	unsignPageIndex := uint64(pageIndex)
	if err != nil {
		fmt.Println("searchKey: ", searchKey, " requestUrl: ", requestUrl, " pageIndex: ", pageIndex)
		log.Fatal("~~~~~`", err)
	}
	unsignPageIndex -= 1
	unsignPageIndex = unsignPageIndex * 15
	requestUrl = fmt.Sprintf("%s&pn=%d", requestUrl, unsignPageIndex)
	c := generate.NewFetcher()
	c.OnHTML(engine.parseRule, func(element *colly.HTMLElement) {
		group.Add(1)
		go engine.extractData(element, group)
	})
	fmt.Println("######### requestUrl: ", requestUrl)
	err = c.Visit(requestUrl)
	if err != nil {
		fmt.Println(err)
	}
}
func (engine *BaiduSearchEngine) extractData(element *colly.HTMLElement, group *sync.WaitGroup) {
	defer group.Done()
	href := element.Attr("href")
	title := element.Text
	c := generate.NewFetcher()
	c.OnResponse(func(response *colly.Response) {
		realURL := response.Request.URL.String()
		fmt.Println("realURL : ", realURL)
		isContainBaidu := strings.Contains(realURL, "baidu")
		if isContainBaidu {
			return
		}
		host := response.Request.URL.Host

		_, ok := config.RuleConfig.IgnoreDomain[host]
		if ok {
			return
		}
		isParse := engine.checkIsParse(host)
		if !isParse {
			return
		}
		result := &model.SearchResult{
			Href:    realURL,
			Title:   title,
			IsParse: generate.Ternary(isParse, int64(1), int64(0)).(int64),
			Host:    host,
		}
		engine.parseResultFunc(result)
	})
	err := c.Visit(href)
	if err != nil {
		fmt.Println("-------", err)
	}
}

func (engine *BaiduSearchEngine) checkIsParse(host string) bool {
	for key := range config.RuleConfig.Rules {
		if host == key {
			return true
		}
	}
	return false
}
