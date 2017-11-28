package command

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/usecase/marketplace/command/mocks"
	"github.com/stretchr/testify/assert"
	"testing"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/entity"
)

func TestCreateBidOrderHandlerHandle_ValidCommandGiven_BidOrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := CreateBidOrder{
		ID:      "TestBidOrder",
		BuyerID: "TestBuyer",
		Price:   555,
		Slot: ds.Slot{
			SupplierRating: 0,
			BuyerRating:    0,
			Resources: ds.Resources{
				CpuCores: 4,
				RamBytes: 100000000,
				Storage:  1000000000,
			},
		},
	}

	expectedOrder, _ := entity.NewBidOrder(cmd.ID, cmd.BuyerID, cmd.Price, cmd.Slot)

	storage := mocks.NewMockCreateBidOrderStorage(ctrl)
	storage.EXPECT().Add(&expectedOrder.Order).Times(1).Return(nil)

	h := NewCreateBidOrderHandler(storage)

	// act
	err := h.Handle(cmd)

	// assert
	assert.NoError(t, err)
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
