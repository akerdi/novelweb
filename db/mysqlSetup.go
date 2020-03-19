package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DSN = "root:123456@tcp(localhost:3306)/test?charset=utf&parseTime=True&loc=Local"
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
}
func CloseDB() {
	db.Close()
}
