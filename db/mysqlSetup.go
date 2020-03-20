package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DSN = "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
const DRIVER = "mysql"

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
func OpenDB() {
	var err error
	db, err = gorm.Open(DRIVER, DSN)
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&NovelNet{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&NovelNet{}).Error; err != nil {
			panic(err)
		}
	}
}
func CloseDB() {
	db.Close()
}
