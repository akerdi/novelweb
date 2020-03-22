package schema

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type NovelNet struct {
	Id int `gorm:"primary_key" json:"id"`
	IsParse uint64 `json:"isParse"` // 是否解析过数据，已经存储在redis 中
	URL string `gorm:"type:varchar(50);not null;" json:"url"` // 真实路由
	Title string `gorm:"type:varchar(50);not null;" json:"title"` // 网站抬头
	MD5 string `gorm:"type:varchar(50); not null;unique_index;" json:"md5"` // 使用RealURL+URLTitle 的编码
}

type NovelChapter struct {
	Id int `gorm:"primary_key" json:"id"`
	MD5 string `gorm:"type:varchar(50);not null;unique_index" json:"md5"`
	Name string `gorm:"type.varchar(30);" json:"name"`
	Chapters ChapterElements `gorm:"type:LONGTEXT;not null"`
	LinkPrefix string `gorm:"type.varchar(8);not null;"`
	OriginURL string `gorm:"type.varchar(50);"`
	Domain string `gorm:"type.varchar(20)"`
}

type NovelChapterElement struct {
	Name string `json:"name"` // 章节名称
	Href string `json:"href"` // 章节链接
}

type ChapterElements []NovelChapterElement

//Value json Marshal to byte
func (a ChapterElements) Value() (driver.Value, error) {
	bytes, err := json.Marshal(a)
	return string(bytes), err
}

//Scan string or byte Unmarshal to json
func (a *ChapterElements) Scan(src interface{}) error {
	switch value := src.(type) {
	case string:
		return json.Unmarshal([]byte(value), a)
	case []byte:
		return json.Unmarshal(value, a)
	default:
		return errors.New("not supported")
	}
}

//Value json Marshal to byte
//func (m Map) Value() (driver.Value, error) {
//	bytes, err := json.Marshal(m)
//	return string(bytes), err
//}