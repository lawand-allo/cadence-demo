package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExpectedRoute struct {
	Path   string
	Method string
}

func TestSetupRoutes(t *testing.T) {
	router := gin.New()
	handler := &Handler{
		CreateOrderUsecase: nil,
		GetOrderUsecase:    nil,
		router:             router,
	}
	handler.setupRoutes()
	routes := handler.router.Routes()
	expectedRoutes := map[ExpectedRoute]bool{
		{"/health", "GET"}:         false,
		{"/", "GET"}:               false,
		{"/order/:orderId", "GET"}: false,
		{"/order/", "POST"}:        false,
	}

	assert.Len(t, routes, len(expectedRoutes))

	for _, route := range routes {
		expectedRoute := ExpectedRoute{
			Path:   route.Path,
			Method: route.Method,
		}
		if _, ok := expectedRoutes[expectedRoute]; ok {
			expectedRoutes[expectedRoute] = true
		}
	}

	for _, found := range expectedRoutes {
		assert.True(t, found)
	}
}
