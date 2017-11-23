package command

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/marketplace/command/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOrderHandlerHandle_ValidCommandGiven_AskOrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedOrder, _ := entity.NewAskOrder("TestAskOrder", "TestBuyer", 555)

	storage := mocks.NewMockCreateOrderStorage(ctrl)
	storage.EXPECT().Store(expectedOrder).Times(1).Return(nil)

	h := NewCreateOrderHandler(storage)

	// act
	err := h.Handle(CreateOrder{
		ID:        "TestAskOrder",
		BuyerID:   "TestBuyer",
		OrderType: int(entity.ASK),
		Price:     555,
	})

	// assert
	assert.NoError(t, err)
}

func TestCreateOrderHandlerHandle_ValidCommandGiven_BidOrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedOrder, _ := entity.NewBidOrder("TestBidOrder", "TestSupplier", 555)

	storage := mocks.NewMockCreateOrderStorage(ctrl)
	storage.EXPECT().Store(expectedOrder).Times(1).Return(nil)

	h := NewCreateOrderHandler(storage)

	// act
	err := h.Handle(CreateOrder{
		ID:         "TestBidOrder",
		SupplierID: "TestSupplier",
		OrderType:  int(entity.BID),
		Price:      555,
	})

	// assert
	assert.NoError(t, err)
}

func TestCreateOrderHandlerHandle_InValidOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockCreateOrderStorage(ctrl)
	h := NewCreateOrderHandler(storage)

	// act
	err := h.Handle(CreateOrder{
		ID:    "TestOrder",
		Price: 555,
	})

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid order type given %v", 0))
}

func TestCreateOrderHandlerHandle_IncorrectCommandGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockCreateOrderStorage(ctrl)
	h := NewCreateOrderHandler(storage)

	// act
	cmd := unknownCommand{}
	err := h.Handle(cmd)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid command %v given", cmd))
}
