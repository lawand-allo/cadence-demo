package usecase

import (
	"cadence-demo/model"
	"cadence-demo/usecase/dependency"
	"github.com/google/uuid"
)

type GetOrderUsecase struct {
	repository dependency.OrderRepository
}

func NewGetOrderUsecase(repository dependency.OrderRepository) *GetOrderUsecase {
	return &GetOrderUsecase{
		repository: repository,
	}
}

func (uc *GetOrderUsecase) GetOrder(orderId uuid.UUID) (*model.GetOrderResponse, error) {
	order, err := uc.repository.ReadOrder(orderId)
	if err != nil {
		return nil, err
	}
	response := &model.GetOrderResponse{
		OrderId: orderId,
		Order:   *order,
	}
	return response, nil
}
