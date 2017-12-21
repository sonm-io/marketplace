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

func TestCreateAskOrderHandlerHandle_ValidCommandGiven_BidOrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := CreateAskOrder{
		ID:         "TestAdkOrder",
		SupplierID: "TestSupplier",
		Price:      "555",
		Slot: ds.Slot{
			SupplierRating: 0,
			BuyerRating:    0,
			Resources: ds.Resources{
				CPUCores: 4,
				RAMBytes: 100000000,
				Storage:  1000000000,
			},
		},
	}

	expectedOrder, _ := entity.NewAskOrder(cmd.ID, cmd.SupplierID, cmd.BuyerID, cmd.Price, cmd.Slot)

	storage := mocks.NewMockCreateAskOrderStorage(ctrl)
	storage.EXPECT().Add(&expectedOrder.Order).Times(1).Return(nil)

	h := NewCreateAskOrderHandler(storage)

	// act
	err := h.Handle(cmd)

	// assert
	assert.NoError(t, err)
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
