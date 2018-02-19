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

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/sonm-io/marketplace/interface/grpc/srv/mocks"
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

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().CreateBidOrder(*req).Times(1)

	serviceMock.EXPECT().OrderByID(req.Id, &pb.Order{})
	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceCreateOrder_ValidAskOrderWithNoResourcesGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	supplierID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType:      pb.OrderType_ASK,
		SupplierID:     supplierID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			SupplierRating: 555,
		},
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().CreateAskOrder(*req).Times(1)

	serviceMock.EXPECT().OrderByID(req.Id, &pb.Order{})

	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(supplierID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceCreateOrder_OrderWithSingleGPUGiven_ErrorReturned(t *testing.T) {
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
				GpuCount: pb.GPUCount_SINGLE_GPU,
			},
		},
	}

	m := NewMarketplace(mocks.NewMockMarketService(ctrl))
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = SINGLE_GPU has been deprecated, "+
		"only NO_GPU and MULTIPLE_GPU are allowed")
}

func TestMarketplaceCreateOrder_OrderWithNoIDGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	expectedID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
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

	order := *req
	IDGenerator = func() string {
		return expectedID
	}
	order.Id = expectedID

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().CreateBidOrder(order).Times(1)

	serviceMock.EXPECT().OrderByID(order.Id, &pb.Order{}).Times(1)

	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceCreateOrder_AnErrorOccurredWhileCreatingOrder_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	supplierID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType:      pb.OrderType_ASK,
		SupplierID:     supplierID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			SupplierRating: 555,
		},
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().CreateAskOrder(*req).Return(fmt.Errorf("some error"))

	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(supplierID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = cannot create order: some error")
}

func TestMarketplaceCreateOrder_AnErrorOccurredWhileRetrievingCreatedOrder_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	supplierID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType:      pb.OrderType_ASK,
		SupplierID:     supplierID,
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			SupplierRating: 555,
		},
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().CreateAskOrder(*req).Times(1)

	serviceMock.EXPECT().OrderByID(req.Id, &pb.Order{}).
		Return(fmt.Errorf("some error"))

	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(supplierID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.EqualError(t, err, "rpc error: code = Internal desc = cannot get order: some error")
}

func TestMarketplaceCreateOrder_InvalidOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID:        buyerID,
		PricePerSecond: pricePerSecond,
	}

	m := NewMarketplace(nil)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err = m.CreateOrder(ctx, req)

	// assert
	assert.EqualError(t, err,
		"rpc error: code = InvalidArgument desc = auth failed: incorrect order type given: ANY")
}
