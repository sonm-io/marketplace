package command

import (
	"testing"

	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/usecase/marketplace/command/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCancelOrderHandlerHandle_ExistingIDGiven_OrderCanceled(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderCanceler(ctrl)
	storage.EXPECT().Cancel("test_order").Times(1).Return(nil)

	h := NewCancelOrderHandler(storage)

	// act
	err := h.Handle(CancelOrder{ID: "test_order"})

	// assert
	assert.NoError(t, err)
}

func TestCancelOrderHandlerHandle_IncorrectCommandGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderCanceler(ctrl)
	h := NewCancelOrderHandler(storage)

	// act
	cmd := unknownCommand{}
	err := h.Handle(cmd)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid command %v given", cmd))
}
