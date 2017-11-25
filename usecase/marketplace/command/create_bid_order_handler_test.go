package command

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/marketplace/command/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBidOrderHandlerHandle_ValidCommandGiven_BidOrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := CreateBidOrder{
		ID:        "TestBidOrder",
		BuyerID:   "TestBuyer",
		OrderType: int(entity.BID),
		Price:     555,
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

	expectedOrder, _ := newBidOrder(cmd)

	storage := mocks.NewMockCreateBidOrderStorage(ctrl)
	storage.EXPECT().Store(expectedOrder).Times(1).Return(nil)

	h := NewCreateBidOrderHandler(storage)

	// act
	err := h.Handle(cmd)

	// assert
	assert.NoError(t, err)
}

func TestCreateBidOrderHandlerHandle_InValidOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockCreateBidOrderStorage(ctrl)
	h := NewCreateBidOrderHandler(storage)

	// act
	err := h.Handle(CreateBidOrder{
		ID:    "TestOrder",
		Price: 555,
	})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid order type given: expected bid order, but got %v", 0))
}

func TestCreateBidOrderHandlerHandle_IncorrectCommandGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockCreateBidOrderStorage(ctrl)
	h := NewCreateBidOrderHandler(storage)

	// act
	cmd := unknownCommand{}
	err := h.Handle(cmd)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid command %v given", cmd))
}
