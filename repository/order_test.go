package repository

import (
	"cadence-demo/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepository(t *testing.T) {
	repository := NewRepository()
	assert.Equal(t, &Repository{storedOrders: map[uuid.UUID]*model.Order{}}, repository)
}
func TestSaveOrder(t *testing.T) {
	repository := NewRepository()
	someUUID := uuid.New()
	expectedOrder := model.Order{
		OrderName: "pizza margherita",
		State:     "pending",
		PickUp:    true,
	}
	expectedOrdersMap := map[uuid.UUID]*model.Order{}
	expectedOrdersMap[someUUID] = &expectedOrder

	repository.SaveOrder(someUUID, expectedOrder)
	savedOrder := repository.storedOrders[someUUID]

	assert.Equal(t, expectedOrdersMap, repository.storedOrders)
	assert.Equal(t, expectedOrder, *savedOrder)

}

func TestReadOrder(t *testing.T) {
	repository := NewRepository()
	someUUID := uuid.New()
	expectedOrder := &model.Order{
		OrderName: "pizza margherita",
		State:     "completed",
		PickUp:    true,
	}
	someOrdersMap := map[uuid.UUID]*model.Order{}
	someOrdersMap[someUUID] = expectedOrder
	repository.storedOrders = someOrdersMap
	readOrder, err := repository.ReadOrder(someUUID)
	assert.Equal(t, expectedOrder, readOrder)
	assert.NoError(t, err)
}

func TestUpdateOrderState(t *testing.T) {
	repository := NewRepository()
	someUUID := uuid.New()
	testOrder := model.Order{
		OrderName: "pizza margherita",
		State:     "pending",
		PickUp:    false,
	}
	someOrdersMap := map[uuid.UUID]*model.Order{}
	someOrdersMap[someUUID] = &testOrder
	repository.storedOrders = someOrdersMap
	expectedState := "completed"
	repository.UpdateOrderState(someUUID, expectedState)

	readOrder, found := repository.storedOrders[someUUID]
	assert.Equal(t, expectedState, readOrder.State)
	assert.True(t, found)
}
