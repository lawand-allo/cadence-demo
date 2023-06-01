package http

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", gin.HandlerFunc(func(ctx *gin.Context) {
		handleIndex(ctx)
	}))
	router.GET("/health", gin.HandlerFunc(func(ctx *gin.Context) {
		handleHealth(ctx)
	}))

	commonGroup := router.Group("/order")
	{
		commonGroup.POST("/", gin.HandlerFunc(func(ctx *gin.Context) {
			handlePostOrder(ctx)
		}))
		commonGroup.GET("/:tripId", gin.HandlerFunc(func(ctx *gin.Context) {
			handleGetOrder(ctx)
		}))
	}
	return router
}
