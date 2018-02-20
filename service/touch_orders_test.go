package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mds "github.com/sonm-io/marketplace/mapper/datastruct"
	"github.com/sonm-io/marketplace/service/mocks"
)

func TestMarketServiceTouchOrders_ValidIDsGiven_OrderTTLUpdated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"}

	storageMock := mocks.NewMockStorage(ctrl)
	storageMock.EXPECT().
		UpdateRow(`UPDATE "orders" SET "status" = ? WHERE (id IN ?) AND (status != ?)`,
			mds.Active, orderIDs, mds.Cancelled).
		Return(nil)

	m := NewMarketService(storageMock)

	// act
	err := m.TouchOrders(orderIDs)

	// assert
	assert.NoError(t, err)
}

func TestMarketServiceTouchOrders_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderIDs := []string{"1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"}

	storageMock := mocks.NewMockStorage(ctrl)
	storageMock.EXPECT().UpdateRow(
		`UPDATE "orders" SET "status" = ? WHERE (id IN ?) AND (status != ?)`,
		mds.Active, orderIDs, mds.Cancelled).
		Return(errors.New("some error"))

	m := NewMarketService(storageMock)

	// act
	err := m.TouchOrders([]string{"1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"})

	// assert
	assert.EqualError(t, err, "cannot update orders' ttl some error")
}
