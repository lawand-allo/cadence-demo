package http

import (
	"cadence-demo/model"
	"cadence-demo/usecase/order"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

func handleIndex(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func handleHealth(ctx *gin.Context) {
	type output struct {
		Alive bool `json:"alive"`
	}
	ctx.JSON(http.StatusOK, &output{Alive: true})
}

func handleGetOrder(ctx *gin.Context) {
	orderId := ctx.Query("orderId")
	if orderId == "" {
		ctx.Error(errors.New("required path parameter 'orderId' can't be empty"))
		ctx.Status(http.StatusBadRequest)
		return
	}
	orderUUID, parseOrderIdDErr := uuid.Parse(orderId)
	if parseOrderIdDErr != nil {
		ctx.Error(errors.WithMessage(parseOrderIdDErr, "couldn't parse provided 'orderId. Please provide a valid UUID"))
		ctx.Status(http.StatusBadRequest)
		return
	}
	response, getOrderErr := order.GetOrder(ctx.Request.Context(), orderUUID)
	if getOrderErr != nil {
		ctx.Error(errors.WithMessage(getOrderErr, "couldn't retrieve the specified order"))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func handlePostOrder(ctx *gin.Context) {
	var createOrderRequest model.CreateOrderRequest

	parseError := ctx.Bind(&createOrderRequest)
	if parseError != nil {
		ctx.Error(errors.Wrap(parseError, "error parsing the request body. Please check the API specification"))
		return
	}
	orderId, createOrderErr := order.CreateOrder(ctx.Request.Context(), createOrderRequest)
	if createOrderErr != nil {
		ctx.Error(errors.WithMessage(createOrderErr, "couldn't create an order"))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, orderId)
}
