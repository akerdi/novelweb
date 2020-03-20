package db

type NovelNet struct {
	Id int `gorm:"primary_key"`
	IsParse uint64 // 是否解析过数据，已经存储在redis 中
	RealURL string `gorm:"type:varchar(50);not null;"` // 真实路由
	URLTitle string `gorm:"type:varchar(50);not null;"` // 网站抬头
}