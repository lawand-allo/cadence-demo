package usecase

import (
	"cadence-demo/model"
	"cadence-demo/repository"
	"context"
	"github.com/google/uuid"
)

type GetOrderUsecase struct {
	repository *repository.Repository
}

func NewGetOrderUsecase(repository *repository.Repository) *GetOrderUsecase {
	return &GetOrderUsecase{
		repository: repository,
	}
}

func (uc *GetOrderUsecase) GetOrder(ctx context.Context, orderId uuid.UUID) (*model.GetOrderResponse, error) {
	return &model.GetOrderResponse{}, nil
}
