package usecase

import (
	"cadence-demo/model"
	"context"
	"github.com/google/uuid"
)

//go:generate mockgen --build_flags=--mod=mod -destination ./mock/create_order.go -package mock . CreateOrderUsecase

type CreateOrderUsecase interface {
	CreateOrder(ctx context.Context, request model.CreateOrderRequest) (uuid.UUID, error)
}
