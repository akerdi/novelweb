package app

import (
	"fmt"
	"novelweb/config"
	"novelweb/db"
	"novelweb/middlewares"
	"novelweb/routes"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func Init() *gin.Engine {
	initConfig()
	// 先打开数据库
	db.OpenDB()
	
	g := gin.New()
	initMiddleware(g)
	return g
}

func initConfig() {
	config := config.InitConfig()
	fmt.Println("initConfig", config)
}

func initMiddleware(g *gin.Engine) {
	g.StaticFile("/favicon.ico", "./public/static/img/favicon.ico")
	g.Static("/assets", "public/assets")
	g.Use(middlewares.Logger())
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	
	routes.InitRoute(g)
	g.Use(static.Serve("/", static.LocalFile("./public", true)))
}