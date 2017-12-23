package srv

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sonm-io/marketplace/infra/grpc/interceptor"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/intf/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestMarketplaceCancelOrder_ValidBidOrderGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	orderID := "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"

	q := query.GetOrder{ID: orderID}
	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:        orderID,
			OrderType: ds.Bid,
			BuyerID:   buyerID,
		},
	}

	orderByIDMock := mocks.NewMockQueryHandler(ctrl)
	orderByIDMock.EXPECT().Handle(q, &report.GetOrderReport{}).
		SetArg(1, orderReport).
		Return(nil)

	cancelOrderMock := mocks.NewMockCommandHandler(ctrl)
	cancelOrderMock.EXPECT().Handle(command.CancelOrder{ID: orderID}).Return(nil)

	m := NewMarketplace(cancelOrderMock, orderByIDMock, nil)
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
	q := query.GetOrder{ID: orderID}

	orderByIDMock := mocks.NewMockQueryHandler(ctrl)
	orderByIDMock.EXPECT().
		Handle(q, &report.GetOrderReport{}).
		Return(fmt.Errorf("order is not found"))

	m := NewMarketplace(nil, orderByIDMock, nil)

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

	q := query.GetOrder{ID: orderID}
	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:        orderID,
			OrderType: ds.Bid,
			BuyerID:   buyerID,
		},
	}

	orderByIDMock := mocks.NewMockQueryHandler(ctrl)
	orderByIDMock.EXPECT().Handle(q, &report.GetOrderReport{}).
		SetArg(1, orderReport).
		Return(nil)

	cancelOrderMock := mocks.NewMockCommandHandler(ctrl)
	cancelOrderMock.EXPECT().
		Handle(command.CancelOrder{ID: orderID}).
		Return(fmt.Errorf("cannot cancel order"))

	m := NewMarketplace(cancelOrderMock, orderByIDMock, nil)
	ctx := interceptor.EthAddrToContext(context.Background(), common.HexToAddress(buyerID))

	// act
	_, err := m.CancelOrder(ctx, &pb.Order{Id: orderID})

	// assert
	assert.EqualError(t, err, "cannot cancel order")
}
