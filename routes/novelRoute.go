package routes

import (
	"novelweb/service"

	"github.com/gin-gonic/gin"
)

var novelRoutes = []Route{
	{
		Path:   "/api/novel/search/:keyword/:page",
		Method: "GET",
		Middles: []gin.HandlerFunc{
			service.SearchNovel(),
		},
	},
	{
		Path:   "/api/novel/chapter/:md5",
		Method: "GET",
		Middles: []gin.HandlerFunc{
			service.SearchChapter(),
		},
	},
	{
		Path:		"/api/novel/content/:md5/:chapterIndex",
		Method: "GET",
		Middles:	[]gin.HandlerFunc{
			service.SearchContent(),
		},
	},
}
