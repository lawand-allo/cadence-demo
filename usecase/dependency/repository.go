package dependency

import (
	"cadence-demo/model"
	"github.com/google/uuid"
)

//go:generate mockgen --build_flags=--mod=mod -destination ./mock/repository.go -package mock . OrderRepository

type OrderRepository interface {
	ReadOrder(orderId uuid.UUID) (*model.Order, error)
	SaveOrder(orderId uuid.UUID, order model.Order)
	UpdateOrderState(orderId uuid.UUID, state model.OrderState)
}
