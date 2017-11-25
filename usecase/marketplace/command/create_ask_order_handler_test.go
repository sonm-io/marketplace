package command

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/marketplace/command/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAskOrderHandlerHandle_ValidCommandGiven_BidOrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedOrder, _ := entity.NewAskOrder("TestAskOrder", "TestSupplier", 555, entity.Slot{})

	storage := mocks.NewMockCreateAskOrderStorage(ctrl)
	storage.EXPECT().Store(expectedOrder).Times(1).Return(nil)

	h := NewCreateAskOrderHandler(storage)

	// act
	err := h.Handle(CreateAskOrder{
		ID:         "TestAskOrder",
		SupplierID: "TestSupplier",
		OrderType:  int(entity.BID),
		Price:      555,
	})

	// assert
	assert.NoError(t, err)
}

func TestCreateAskOrderHandlerHandle_InValidOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockCreateAskOrderStorage(ctrl)
	h := NewCreateAskOrderHandler(storage)

	// act
	err := h.Handle(CreateAskOrder{
		ID:    "TestOrder",
		Price: 555,
	})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid order type given: expected bid order, but got %v", 0))
}

func TestCreateAskOrderHandlerHandle_IncorrectCommandGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockCreateAskOrderStorage(ctrl)
	h := NewCreateAskOrderHandler(storage)

	// act
	cmd := unknownCommand{}
	err := h.Handle(cmd)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid command %v given", cmd))
}
