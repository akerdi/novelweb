package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

/*
Config.Rule
	link_prefix: ||| -1 代表取chapter_url ||| 1 代表直接取URL ||| 0 代表使用域名加拼接
	chapter_selector: 用于寻找chapter 目录章节的元素
	content_selector: 用于寻找content 的元素
	chapter_tail: 如果存在，则作为chapter 附加添加到link 的后缀
*/

type ruleConfig struct {
	Engines []string
	Rules map[string]map[string]interface{} `json:"rules"`
	IgnoreDomain map[string]int	 `json:"ignores"`
}

var RuleConfig *ruleConfig

func init() {
	// link_prefix ||| -1 代表 章节页面网址+章节a.href ||| 0 代表使用域名+章节a.href ||| 1 代表章节a.href ||| 2 代表需要拼接后缀.html（如/dir.html 等, 如果已存在.html 则不再添加）
	RuleConfig = &ruleConfig{}
	RuleConfig.Engines = []string{"baidu"}
	ConfigRule()
}

func ConfigRule() {
	fileReader, err := os.Open("rule/rule.json")
	if err != nil {
		log.Fatal("read rule.json meet error: ", err)
	}
	b, _ := ioutil.ReadAll(fileReader)

	var rule ruleConfig
	err = json.Unmarshal(b, &rule)
	if err != nil {
		log.Fatal("unmarshal rule meet error: ", err)
	}
	RuleConfig.Rules = rule.Rules
	RuleConfig.IgnoreDomain = rule.IgnoreDomain
}
