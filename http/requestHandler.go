package http

import (
	usecase "cadence-demo/http/dependency/usecase"
	"cadence-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"net/http"
)

type Handler struct {
	CreateOrderUsecase usecase.CreateOrderUsecase
	GetOrderUsecase    usecase.GetOrderUsecase
	router             *gin.Engine
}

func NewHandler(createOrderUsecase usecase.CreateOrderUsecase, getOrderUsecase usecase.GetOrderUsecase) *Handler {
	router := gin.New()
	handler := &Handler{
		CreateOrderUsecase: createOrderUsecase,
		GetOrderUsecase:    getOrderUsecase,
		router:             router,
	}
	handler.setupRoutes()
	return handler
}

func (h *Handler) handleIndex(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (h *Handler) handleHealth(ctx *gin.Context) {
	type output struct {
		Alive bool `json:"alive"`
	}
	ctx.JSON(http.StatusOK, &output{Alive: true})
}

func (h *Handler) handleGetOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
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
	response, getOrderErr := h.GetOrderUsecase.GetOrder(ctx.Request.Context(), orderUUID)
	if getOrderErr != nil {
		ctx.Error(errors.WithMessage(getOrderErr, "couldn't retrieve the specified order"))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) handlePostOrder(ctx *gin.Context) {
	var createOrderRequest model.CreateOrderRequest

	parseError := ctx.Bind(&createOrderRequest)
	if parseError != nil {
		ctx.Error(errors.Wrap(parseError, "error parsing the request body. Please check the API specification"))
		return
	}
	orderId, createOrderErr := h.CreateOrderUsecase.CreateOrder(ctx.Request.Context(), createOrderRequest)
	if createOrderErr != nil {
		ctx.Error(errors.WithMessage(createOrderErr, "couldn't create an order"))
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, orderId)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
