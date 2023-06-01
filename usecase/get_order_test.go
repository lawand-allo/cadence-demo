package usecase

import (
	"cadence-demo/model"
	"cadence-demo/usecase/dependency/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOrder(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mock.NewMockOrderRepository(mockController)
	getOrderUseCase := NewGetOrderUsecase(mockRepository)
	someUUID := uuid.New()
	someOrder := &model.Order{
		OrderName: "burger",
		State:     model.Completed,
		PickUp:    true,
	}
	mockRepository.EXPECT().ReadOrder(someUUID).Times(1).Return(someOrder, nil)
	expectedResponse := &model.GetOrderResponse{
		OrderId: someUUID,
		Order:   *someOrder,
	}

	response, err := getOrderUseCase.GetOrder(someUUID)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestGetOrderNotFound(t *testing.T) {
	mockController := gomock.NewController(t)
	mockRepository := mock.NewMockOrderRepository(mockController)
	getOrderUseCase := NewGetOrderUsecase(mockRepository)
	someUUID := uuid.New()
	expectedError := errors.New("order not found")
	mockRepository.EXPECT().ReadOrder(someUUID).Times(1).Return(nil, expectedError)

	_, err := getOrderUseCase.GetOrder(someUUID)
	assert.ErrorContains(t, err, expectedError.Error())
}
