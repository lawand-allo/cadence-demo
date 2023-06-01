package usecase

import (
	"cadence-demo/model"
	"cadence-demo/usecase/dependency"
	"context"
	"github.com/google/uuid"
)

type CreateOrderUsecase struct {
	repository dependency.OrderRepository
}

func NewCreateOrderUsecase(repository dependency.OrderRepository) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		repository: repository,
	}
}

func (uc *CreateOrderUsecase) CreateOrder(ctx context.Context, request model.CreateOrderRequest) (uuid.UUID, error) {
	return uuid.New(), nil

	// save order to repo
	// start do async processing
	// 		update order state
}
