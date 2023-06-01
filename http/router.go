package http

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) setupRoutes() {

	h.router.GET("/", gin.HandlerFunc(func(ctx *gin.Context) {
		h.handleIndex(ctx)
	}))
	h.router.GET("/health", gin.HandlerFunc(func(ctx *gin.Context) {
		h.handleHealth(ctx)
	}))

	commonGroup := h.router.Group("/order")
	{
		commonGroup.POST("/", gin.HandlerFunc(func(ctx *gin.Context) {
			h.handlePostOrder(ctx)
		}))
		commonGroup.GET("/:orderId", gin.HandlerFunc(func(ctx *gin.Context) {
			h.handleGetOrder(ctx)
		}))
	}
}
