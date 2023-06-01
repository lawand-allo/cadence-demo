package repository

import (
	"cadence-demo/model"
	"errors"
	"fmt"
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

func (r *Repository) SaveOrder(orderId uuid.UUID, order model.Order) {
	r.storedOrders[orderId] = &order
}

func (r *Repository) ReadOrder(orderId uuid.UUID) (*model.Order, error) {
	order, ok := r.storedOrders[orderId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("the requested order with ID %v not found", orderId))
	}
	return order, nil
}

func (r *Repository) UpdateOrderState(orderId uuid.UUID, state string) {
	(&*r.storedOrders[orderId]).State = state
}
