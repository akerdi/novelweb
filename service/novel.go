package service

import (
	"fmt"
	"net/http"

	"novelweb/db"

	"github.com/gin-gonic/gin"
)

func SearchNovel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("~~~~~~~~~~~~~~~~")
		keyWord := c.Param("keyword")
		// 通过在网络查找是否有对应的数据
		// 保存数据至db 中

		dbFD := db.GetDB()
		newNovel := db.NovelNet{
			IsParse:  1,
			RealURL:  "www.baidu.com",
			URLTitle: "baidu",
		}
		fmt.Println("~~~~~~", dbFD.Create(&newNovel))
		c.String(http.StatusOK, fmt.Sprintf("%s%s", "我喜欢吃饭", keyWord))
	}
}
