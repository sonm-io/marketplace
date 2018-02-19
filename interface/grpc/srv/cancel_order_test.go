package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sonm-io/marketplace/infra/grpc/interceptor"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"github.com/sonm-io/marketplace/interface/grpc/srv/mocks"
)

func TestMarketplaceCancelOrder_ValidBidOrderGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	orderID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"

	order := pb.Order{
		Id:        orderID,
		OrderType: pb.OrderType_BID,
		ByuerID:   buyerID,
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().OrderByID(orderID, &pb.Order{}).
		SetArg(1, order).
		Return(nil)

	serviceMock.EXPECT().CancelOrder(orderID).
		Return(nil)

	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err := m.CancelOrder(ctx, &pb.Order{Id: orderID})

	// assert
	assert.NoError(t, err)
}

func TestMarketplaceCancelOrder_InExistentOrderGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"

	orderByIDMock := mocks.NewMockMarketService(ctrl)
	orderByIDMock.EXPECT().
		OrderByID(orderID, &pb.Order{}).
		Return(fmt.Errorf("order is not found"))

	m := NewMarketplace(orderByIDMock)

	// act
	_, err := m.CancelOrder(context.Background(), &pb.Order{Id: orderID})

	// assert
	assert.EqualError(t, err,
		"rpc error: code = Internal desc = cannot cancel order: order is not found")
}

func TestMarketplaceCancelOrder_AnErrorOccurredWhileCancelingOrder_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	orderID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"

	order := pb.Order{
		Id:        orderID,
		OrderType: pb.OrderType_BID,
		ByuerID:   buyerID,
	}

	serviceMock := mocks.NewMockMarketService(ctrl)
	serviceMock.EXPECT().OrderByID(orderID, &pb.Order{}).
		SetArg(1, order).
		Return(nil)

	serviceMock.EXPECT().CancelOrder(orderID).
		Return(fmt.Errorf("cannot cancel order"))

	m := NewMarketplace(serviceMock)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err := m.CancelOrder(ctx, &pb.Order{Id: orderID})

	// assert
	assert.EqualError(t, err, "cannot cancel order")
}
