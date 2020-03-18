package middlewares

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
)

func NotRouteResponse() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Status(404)
	}
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		start := time.Now()
		str := fmt.Sprintf("%s::%s : %s : %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		fmt.Println(str)
		context.Next()
		end := time.Now()
		fmt.Println(context.Writer.Status(), end.Sub(start))
	}
}

func ReturnPublic() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		if method == "GET" {
			fmt.Println("????", "public")
			context.File("./public")
		} else {
			context.Next()
		}
	}
}

func Test() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, "chifanchifan")
	}
}