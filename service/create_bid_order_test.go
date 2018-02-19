package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mds "github.com/sonm-io/marketplace/mapper/datastruct"
	pb "github.com/sonm-io/marketplace/proto"
	"github.com/sonm-io/marketplace/service/mocks"
)

func TestMarketServiceCreateBidOrder_IncorrectOrderGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := mocks.NewMockStorage(ctrl)
	m := NewMarketService(storageMock)

	// act
	err := m.CreateBidOrder(pb.Order{})

	// assert
	assert.EqualError(t, err, "price cannot be nil")
}

func TestMarketServiceCreateBidOrder_EmptyBuyerIDGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := mocks.NewMockStorage(ctrl)
	m := NewMarketService(storageMock)

	order := pb.Order{
		OrderType:      pb.OrderType_BID,
		PricePerSecond: pb.NewBigIntFromInt(555),
	}

	// act
	err := m.CreateBidOrder(order)

	// assert
	assert.EqualError(t, err, "buyer is required")
}

func TestMarketServiceCreateBidOrder_StorageErrorOccurred_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageMock := mocks.NewMockStorage(ctrl)
	storageMock.EXPECT().InsertRow(gomock.Any(), gomock.Any()).Return(errors.New("some error"))

	m := NewMarketService(storageMock)

	order := pb.Order{
		OrderType:      pb.OrderType_BID,
		ByuerID:        "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159",
		PricePerSecond: pb.NewBigIntFromInt(555),
	}

	// act
	err := m.CreateBidOrder(order)

	// assert
	assert.EqualError(t, err, "cannot create a new order: some error")
}

func TestMarketServiceCreateBidOrder_ValidOrderGiven_OrderCreated(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	order := pb.Order{
		Id:             "1b5dfa00-af3c-4e2d-b64b-c5d62e89430b",
		OrderType:      pb.OrderType_BID,
		ByuerID:        "0x9B27D3C3571731deDb23EaFEa34a3a6E05daE159",
		PricePerSecond: pb.NewBigIntFromInt(555),
	}

	query := `INSERT INTO "orders" ("id","type","supplier_id","buyer_id","price","slot_duration","slot_buyer_rating",` +
		`"slot_supplier_rating","resources_cpu_cores","resources_ram_bytes","resources_gpu_count","resources_storage",` +
		`"resources_net_inbound","resources_net_outbound","resources_net_type","resources_properties","status") ` +
		`VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	storageMock := mocks.NewMockStorage(ctrl)
	storageMock.EXPECT().InsertRow(query,
		order.Id, order.OrderType, "", order.ByuerID, order.PricePerSecond.Unwrap().String(),
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, mds.Active,
	)

	m := NewMarketService(storageMock)

	// act
	err := m.CreateBidOrder(order)

	// assert
	assert.NoError(t, err)
}
