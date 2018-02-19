package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mds "github.com/sonm-io/marketplace/mapper/datastruct"
	"github.com/sonm-io/marketplace/service/mocks"
)

func TestMarketServiceCancelOrder_ValidIDGiven_OrderCancelled(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderID := "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b"

	storageMock := mocks.NewMockStorage(ctrl)
	storageMock.EXPECT().
		UpdateRow(`UPDATE "orders" SET "status" = ? WHERE (id = ?)`, mds.Cancelled, orderID).
		Return(nil)

	m := NewMarketService(storageMock)

	// act
	err := m.CancelOrder(orderID)

	// assert
	assert.NoError(t, err)
}

func TestMarketServiceCancelOrder_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := mocks.NewMockStorage(ctrl)
	storageMock.EXPECT().UpdateRow(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

	m := NewMarketService(storageMock)

	// act
	err := m.CancelOrder("1b5dfa00-af3c-4e2d-b64b-c5d62e89430b")

	// assert
	assert.EqualError(t, err, "cannot cancel order: some error")
}
