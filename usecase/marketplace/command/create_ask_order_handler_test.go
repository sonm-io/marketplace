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

	cmd := CreateAskOrder{
		ID:         "TestAdkOrder",
		SupplierID: "TestSupplier",
		OrderType:  int(entity.BID),
		Price:      555,
		Slot: Slot{
			SupplierRating: 0,
			BuyerRating:    0,
			Resources: Resources{
				CpuCores: 4,
				RamBytes: 100000000,
				Storage:  1000000000,
			},
		},
	}

	expectedOrder, _ := newAskOrder(cmd)

	storage := mocks.NewMockCreateAskOrderStorage(ctrl)
	storage.EXPECT().Add(expectedOrder).Times(1).Return(nil)

	h := NewCreateAskOrderHandler(storage)

	// act
	err := h.Handle(cmd)

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
