package storage

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/interface/storage/mocks"
)

func TestOrderStorageByID_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	order := &ds.Order{}
	order.ID = "test_order"

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Get("test_order").Return(order, nil)

	s := NewOrderStorage(engineMock)

	// act
	obtained, err := s.ByID("test_order")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, order, &obtained)
}

func TestOrderStorageAdd_ValidOrderGiven_OrderStored(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := ds.Order{}
	expected.ID = "test_order"

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Add(&expected, "test_order").Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	err := s.Add(&expected)

	// assert
	assert.NoError(t, err, "non-error result expected")
}

func TestOrderStorageRemove_ExistingIDGiven_OrderRemoved(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Remove("test_order").Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	err := s.Remove("test_order")

	// assert
	assert.NoError(t, err, "non-error result expected")
}
