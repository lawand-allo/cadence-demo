package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", gin.HandlerFunc(func(ctx *gin.Context) {
		handleIndex(ctx)
	}))
	router.GET("/health", gin.HandlerFunc(func(ctx *gin.Context) {
		handleHealth(ctx)
	}))

	return router
}

func handleIndex(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func handleHealth(ctx *gin.Context) {
	type output struct {
		Alive bool `json:"alive"`
	}
	ctx.JSON(http.StatusOK, &output{Alive: true})
}
