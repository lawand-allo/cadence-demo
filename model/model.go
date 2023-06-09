package model

import "github.com/google/uuid"

type CreateOrderRequest struct {
	OrderName string `json:"name"`
	PickUp    bool   `json:"pickUp,omitempty""`
}

type GetOrderResponse struct {
	OrderId uuid.UUID `json:"orderId" binding:"required"`
	Order
}

type Order struct {
	OrderName string     `json:"name" binding:"required"`
	State     OrderState `json:"state" binding:"required"`
	PickUp    bool       `json:"pickUp" binding:"required"`
}

type OrderState string

const (
	Pending   OrderState = "PENDING"
	Running   OrderState = "RUNNING"
	Completed OrderState = "COMPLETED"
	Failed    OrderState = "FAILED"
)
