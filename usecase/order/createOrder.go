package order

import (
	"cadence-demo/model"
	"context"
	"github.com/google/uuid"
)

func CreateOrder(ctx context.Context, request model.CreateOrderRequest) (uuid.UUID, error) {
	return uuid.New(), nil
}
