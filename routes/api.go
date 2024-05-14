package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api(ctx *gin.Engine) {
	api := ctx.Group("/api")

	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	// CDN group
	cdn := api.Group("/cdn")
	cdn.GET("/image/:filename", func(ctx *gin.Context) {})
	cdn.GET("/doc/:filename", func(ctx *gin.Context) {})
	
	// Upload group
	upload := api.Group("/upload")
	upload.POST("/image", func(ctx *gin.Context) {})
	upload.POST("/doc", func(ctx *gin.Context) {})
}