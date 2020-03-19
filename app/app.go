package app

import (
	"fmt"
	"novelweb/middlewares"
	"novelweb/config"
	"novelweb/routes"
	"novelweb/db"
	
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
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
	g.StaticFile("/favicon.ico", "./resource/favicon.ico")
	g.Static("/assets", "public/assets")
	g.Use(middlewares.Logger())
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	
	routes.InitRoute(g)
}