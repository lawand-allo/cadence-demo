package usecase

import (
	"cadence-demo/model"
	"github.com/google/uuid"
)

//go:generate mockgen --build_flags=--mod=mod -destination ./mock/get_order.go -package mock . GetOrderUsecase

type GetOrderUsecase interface {
	GetOrder(orderId uuid.UUID) (*model.GetOrderResponse, error)
}
