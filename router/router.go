package router

import (
	"github.com/gin-gonic/gin"
	"hummer/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/hummer/api")
	{
		api.POST("/shorten", handler.Shorten)
	}

	r.GET("/go/:short_code", handler.Redirect)

	return r
}
