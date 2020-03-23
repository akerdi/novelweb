package routes

import (
	"novelweb/service"

	"github.com/gin-gonic/gin"
)

var novelRoutes = []Route{
	{
		Path:   "/api/novel/search/:keyword/:page",
		Method: "POST",
		Middles: []gin.HandlerFunc{
			service.SearchNovel(),
		},
	},
	{
		Path:   "/api/novel/chapter/:md5",
		Method: "POST",
		Middles: []gin.HandlerFunc{
			service.SearchChapter(),
		},
	},
	{
		Path:		"/api/novel/content/:md5/:chapterIndex",
		Method: "POST",
		Middles:	[]gin.HandlerFunc{
			service.SearchContent(),
		},
	},
}
