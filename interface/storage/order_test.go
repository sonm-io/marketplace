package storage

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"
	"github.com/sonm-io/marketplace/interface/storage/mocks"
	"github.com/sonm-io/marketplace/report"
)

func TestOrderStorageByID_ExistingIDGiven_OrderReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := report.GetOrderReport{ID: "test_order"}
	order := entity.Order{ID: "test_order"}

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().ByID("test_order", &entity.Order{}).SetArg(1, order).Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	obtained, err := s.ByID("test_order")

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, expected, obtained)
}

func TestOrderStorageBySpecWithLimit_ValidSpecGiven_OrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := report.GetOrdersReport{
		{
			ID:    "test_obj_101",
			Price: 101,
		},
		{
			ID:    "test_obj_105",
			Price: 105,
		},
	}

	var orders report.GetOrdersReport
	spec := priceIsBetweenTestSpec{From: 101, To: 106}
	q := inmemory.ConcreteCriteria{
		Limit: 10,
		Spec:  spec,
	}

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Match(q, &orders).SetArg(1, expected).Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	obtained, err := s.BySpecWithLimit(spec, 10)

	// assert
	assert.NoError(t, err, "non-error result expected")
	assert.Equal(t, expected, obtained)

}

func TestOrderStorageStore_ValidOrderGiven_OrderStored(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := entity.Order{ID: "test_order"}

	engineMock := mocks.NewMockEngine(ctrl)
	engineMock.EXPECT().Store(&expected).Return(nil)

	s := NewOrderStorage(engineMock)

	// act
	err := s.Store(&expected)

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

type priceIsBetweenTestSpec struct {
	From int64
	To   int64
}

func (s priceIsBetweenTestSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*entity.Order)
	return order.Price >= s.From && order.Price < s.To
}
