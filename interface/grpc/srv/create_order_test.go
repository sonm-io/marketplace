package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sonm-io/marketplace/infra/grpc/interceptor"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/intf/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestMarketplaceCreateOrder_ValidBidOrderGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType:      pb.OrderType_BID,
		ByuerID:        buyerID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	q := query.GetOrder{
		ID: req.GetId(),
	}

	expected := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID:        buyerID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
			BuyerID:        buyerID,
			PricePerSecond: pricePerSecond.Unwrap().String(),
			Slot: &ds.Slot{
				Resources: ds.Resources{
					CPUCores: 4,
					RAMBytes: 10000,
				},
			},
		},
	}

	var (
		cmd command.CreateBidOrder
		c   intf.Command
	)
	bindCreateBidOrderCommand(req, &cmd)
	c = cmd

	createOrderMock := mocks.NewMockCommandHandler(ctrl)
	createOrderMock.EXPECT().Handle(c).
		Return(nil)

	orderByIDMock := mocks.NewMockQueryHandler(ctrl)
	orderByIDMock.EXPECT().Handle(q, &report.GetOrderReport{}).
		SetArg(1, orderReport).
		Return(nil)

	m := NewMarketplace(createOrderMock, orderByIDMock, nil)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	obtained, err := m.CreateOrder(ctx, req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMarketplaceCreateOrder_ValidAskOrderWithNoResourcesGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	supplierID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		OrderType:      pb.OrderType_ASK,
		SupplierID:     supplierID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			SupplierRating: 555,
		},
	}

	expectedID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"
	IDGenerator = func() string {
		return expectedID
	}

	q := query.GetOrder{
		ID: expectedID,
	}

	expected := &pb.Order{
		Id:             expectedID,
		SupplierID:     supplierID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			SupplierRating: 555,
			Resources:      &pb.Resources{},
		},
	}

	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:             expectedID,
			SupplierID:     supplierID,
			PricePerSecond: pricePerSecond.Unwrap().String(),
			Slot: &ds.Slot{
				SupplierRating: 555,
			},
		},
	}

	var (
		cmd command.CreateAskOrder
		c   intf.Command
	)
	bindCreateAskOrderCommand(req, &cmd)
	c = cmd

	createOrderMock := mocks.NewMockCommandHandler(ctrl)
	createOrderMock.EXPECT().Handle(c).
		Return(nil)

	orderByIDMock := mocks.NewMockQueryHandler(ctrl)
	orderByIDMock.EXPECT().Handle(q, &report.GetOrderReport{}).
		SetArg(1, orderReport).
		Return(nil)

	m := NewMarketplace(createOrderMock, orderByIDMock, nil)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(supplierID))

	// act
	obtained, err := m.CreateOrder(ctx, req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMarketplaceCreateOrder_InValidRequest_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:        "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType: pb.OrderType_BID,
		ByuerID:   buyerID,
		Price:     "100",
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	expectedErr := fmt.Errorf("an error occurred")

	var (
		cmd command.CreateBidOrder
		c   intf.Command
	)
	bindCreateBidOrderCommand(req, &cmd)
	c = cmd

	createOrderMock := mocks.NewMockCommandHandler(ctrl)
	createOrderMock.EXPECT().Handle(c).
		Return(expectedErr)

	m := NewMarketplace(createOrderMock, nil, nil)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err := m.CreateOrder(ctx, req)

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = Cannot create order: an error occurred")
}

func TestMarketplaceCreateOrder_InvalidOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID: buyerID,
		Price:   "100",
	}

	m := NewMarketplace(nil, nil, nil)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err := m.CreateOrder(ctx, req)

	// assert
	assert.EqualError(t, err,
		"rpc error: code = InvalidArgument desc = auth failed: incorrect order type given: ANY")
}
