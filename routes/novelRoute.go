package routes

import (
	"novelweb/service"

	"github.com/gin-gonic/gin"
)

var novelRoutes = []Route{
	{
		Path:   "/api/novel/search/:keyword",
		Method: "POST",
		Middles: []gin.HandlerFunc{
			service.SearchNovel(),
		},
	},
}
