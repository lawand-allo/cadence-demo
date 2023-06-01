package repository

import (
	"cadence-demo/model"
	"github.com/google/uuid"
)

type Repository struct {
	storedOrders map[uuid.UUID]*model.Order
}

func NewRepository() *Repository {
	return &Repository{
		storedOrders: map[uuid.UUID]*model.Order{},
	}
}

func (r *Repository) saveOrder(orderId uuid.UUID, order model.Order) {
	r.storedOrders[orderId] = &order
}

func (r *Repository) readOrder(orderId uuid.UUID) *model.Order {
	return r.storedOrders[orderId]
}

func (r *Repository) updateOrderState(orderId uuid.UUID, state string) {
	(&*r.storedOrders[orderId]).State = state
}
