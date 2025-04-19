package router

import (
	"github.com/gin-gonic/gin"
	"hummer/config"
	"hummer/handler"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		println(token, "Bearer "+config.APIToken)
		if token != "Bearer "+config.APIToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api", AuthMiddleware())
	{
		api.POST("/shorten", handler.Shorten)
	}

	r.GET("/:short_code", handler.Redirect)

	return r
}
