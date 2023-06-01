package usecase

import (
	"cadence-demo/model"
	"cadence-demo/usecase/dependency"
	"context"
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

func (uc *GetOrderUsecase) GetOrder(ctx context.Context, orderId uuid.UUID) (*model.GetOrderResponse, error) {
	return &model.GetOrderResponse{}, nil
}
