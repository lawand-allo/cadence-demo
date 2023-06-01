package http

import (
	"bytes"
	"cadence-demo/http/dependency/usecase/mock"
	"cadence-demo/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandleIndex(t *testing.T) {
	mockController := gomock.NewController(t)
	mockGetOrderUsecase := mock.NewMockGetOrderUsecase(mockController)
	mockCreateOrderUsecase := mock.NewMockCreateOrderUsecase(mockController)

	handler := NewHandler(mockCreateOrderUsecase, mockGetOrderUsecase)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	handler.handleIndex(ctx)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestHandleGetHealth(t *testing.T) {
	mockController := gomock.NewController(t)
	mockGetOrderUsecase := mock.NewMockGetOrderUsecase(mockController)
	mockCreateOrderUsecase := mock.NewMockCreateOrderUsecase(mockController)

	handler := NewHandler(mockCreateOrderUsecase, mockGetOrderUsecase)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = &http.Request{
		Method:     "GET",
		URL:        &url.URL{Path: "/health/"},
		RequestURI: "/health/",
	}

	type output struct {
		Alive bool `json:"alive"`
	}

	expectedGetHealthResponse := &output{
		Alive: true,
	}

	handler.handleHealth(ctx)

	var actualResponse output
	json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)
	assert.Equal(t, expectedGetHealthResponse, &actualResponse)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestHandleGetOrder(t *testing.T) {
	mockController := gomock.NewController(t)
	mockGetOrderUsecase := mock.NewMockGetOrderUsecase(mockController)
	mockCreateOrderUsecase := mock.NewMockCreateOrderUsecase(mockController)

	handler := NewHandler(mockCreateOrderUsecase, mockGetOrderUsecase)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	orderUUID := uuid.New()
	orderId := orderUUID.String()
	ctx.AddParam("orderId", orderId)
	ctx.Request = &http.Request{
		Method:     "GET",
		URL:        &url.URL{Path: "/order/" + orderId},
		RequestURI: "/order/" + orderId,
	}
	expectedGetOrderResponse := &model.GetOrderResponse{
		OrderId: orderUUID,
		Order: model.Order{
			OrderName: "pizza",
			State:     "pending",
			PickUp:    false,
		},
	}

	mockGetOrderUsecase.EXPECT().GetOrder(orderUUID).Times(1).Return(expectedGetOrderResponse, nil)

	handler.handleGetOrder(ctx)

	var actualResponse model.GetOrderResponse
	json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	assert.Equal(t, expectedGetOrderResponse, &actualResponse)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestHandlePostOrder(t *testing.T) {
	mockController := gomock.NewController(t)
	mockGetOrderUsecase := mock.NewMockGetOrderUsecase(mockController)
	mockCreateOrderUsecase := mock.NewMockCreateOrderUsecase(mockController)

	handler := NewHandler(mockCreateOrderUsecase, mockGetOrderUsecase)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)

	createOrderRequest := &model.CreateOrderRequest{
		OrderName: "pizza",
		PickUp:    false,
	}
	marshal, _ := json.Marshal(createOrderRequest)
	request, err := http.NewRequest("POST", "/order/", bytes.NewBuffer(marshal))
	if err != nil {
		return
	}
	ctx.Request = request

	expectedOrderUUID := uuid.New()
	expectedOrderId := expectedOrderUUID.String()

	mockCreateOrderUsecase.EXPECT().CreateOrder(ctx.Request.Context(), gomock.Any()).Times(1).Return(expectedOrderUUID, nil)

	handler.handlePostOrder(ctx)

	var actualResponse string
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)
	if err != nil {
		er := err.Error()
		println(er)
		return
	}

	assert.Equal(t, actualResponse, expectedOrderId)
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}
