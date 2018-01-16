package command

import (
	"testing"

	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/sonm-io/marketplace/usecase/marketplace/command/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTouchOrdersHandlerHandle_ValidParamsGiven_OrdersTouched(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}

	storage := mocks.NewMockOrderToucher(ctrl)
	storage.EXPECT().Touch(orderIDs).Times(1).Return(nil)

	h := NewTouchOrdersHandler(storage)

	// act
	err := h.Handle(TouchOrders{IDs: orderIDs})

	// assert
	assert.NoError(t, err)
}

func TestTouchOrdersHandlerHandle_IncorrectCommandGivenErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockOrderToucher(ctrl)
	h := NewTouchOrdersHandler(storage)

	// act
	cmd := unknownCommand{}
	err := h.Handle(cmd)

	// assert
	assert.EqualError(t, err, fmt.Sprintf("invalid command %v given", cmd))
}
