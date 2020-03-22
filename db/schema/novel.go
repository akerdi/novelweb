package schema

type NovelNet struct {
	Id int `gorm:"primary_key" json:"id"`
	IsParse uint64 `json:"isParse"` // 是否解析过数据，已经存储在redis 中
	URL string `gorm:"type:varchar(50);not null;" json:"url"` // 真实路由
	Title string `gorm:"type:varchar(50);not null;" json:"title"` // 网站抬头
	MD5 string `gorm:"type:varchar(50); not null;unique_index;" json:"md5"` // 使用RealURL+URLTitle 的编码
}

type NovelChapter struct {
	Id int `gorm:"primary_key" json:"id"`
	MD5 string `gorm:"type:varchar(50);not null;unique_index;json:"md5"`
	Name string `gorm:"type.varchar(30);" json:"name"`
	Chapters string `gorm:"type.text";`
}