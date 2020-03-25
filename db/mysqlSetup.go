package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"novelweb/config"
	"novelweb/db/schema"
)


const DRIVER = "mysql"

var db *gorm.DB

func GetDB() *gorm.DB {
	return db.Debug()
}
func OpenDB() {
	var err error
	curConfig := config.GetConfig()
	DSN := fmt.Sprintf("root:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", curConfig.DB.Password, curConfig.DB.Host, curConfig.DB.Port, curConfig.DB.Name)
	fmt.Println("DSN: ", DSN)
	db, err = gorm.Open(DRIVER, DSN)
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&schema.NovelNet{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&schema.NovelNet{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&schema.NovelChapter{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&schema.NovelChapter{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&schema.NovelContent{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&schema.NovelContent{}).Error; err != nil {
			panic(err)
		}
	}
}
func CloseDB() {
	db.Close()
}
