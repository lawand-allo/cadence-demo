package model

import "github.com/google/uuid"

type CreateOrderRequest struct {
	OrderName string `json:"name"`
	PickUp    bool   `json:"pickUp,omitempty""`
}

type GetOrderResponse struct {
	OrderId   uuid.UUID `json:"orderId" binding:"required"`
	OrderName string    `json:"name" binding:"required"`
	State     string    `json:"state" binding:"required"`
	PickUp    bool      `json:"pickUp" binding:"required"`
}

type Order struct {
	OrderName string `json:"name" binding:"required"`
	State     string `json:"state" binding:"required"`
	PickUp    bool   `json:"pickUp" binding:"required"`
}
