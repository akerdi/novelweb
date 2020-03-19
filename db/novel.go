package db

type NovelNet struct {
	Id int
	IsParse bool // 是否解析过数据，已经存储在redis 中
	RealURL string // 真实路由
	URLTitle string // 网站抬头
}