package routes

import (
	"log"
	"novelweb/middlewares"
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Method  string
	Middles []gin.HandlerFunc
}

func applyRoutes(g *gin.Engine, routes []Route) {
	for _, route := range routes {
		method := strings.ToUpper(route.Method)
		if method == "GET" {
			g.GET(route.Path, route.Middles...)
		} else if method == "POST" {
			g.POST(route.Path, route.Middles...)
		} else if method == "PUT" {
			g.PUT(route.Path, route.Middles...)
		} else {
			log.Fatal("current not support for method: ", method)
		}
	}
}

func InitRoute(g *gin.Engine) {
	//applyRoutes(g, CommonRoutes)
	applyRoutes(g, novelRoutes)
	g.NoRoute(middlewares.ReturnPublic())
	g.Use(middlewares.NotRouteResponse())
}
