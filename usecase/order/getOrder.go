package order

import (
	"cadence-demo/model"
	"context"
	"github.com/google/uuid"
)

func GetOrder(ctx context.Context, orderId uuid.UUID) (*model.GetOrderResponse, error) {
	return &model.GetOrderResponse{}, nil
}
