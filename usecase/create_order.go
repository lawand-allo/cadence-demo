package usecase

import (
	"cadence-demo/model"
	"cadence-demo/repository"
	"context"
	"github.com/google/uuid"
)

type CreateOrderUsecase struct {
	repository *repository.Repository
}

func NewCreateOrderUsecase(repository *repository.Repository) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		repository: repository,
	}
}

func (uc *CreateOrderUsecase) CreateOrder(ctx context.Context, request model.CreateOrderRequest) (uuid.UUID, error) {
	return uuid.New(), nil
}
