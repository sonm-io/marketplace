package srv

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"fmt"
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

	req := &pb.Order{
		Id:        "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType: pb.OrderType_BID,
		ByuerID:   "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:     "100",
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
		Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:   "100",
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
			BuyerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
			Price:   "100",
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

	// act
	obtained, err := m.CreateOrder(context.Background(), req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMarketplaceCreateOrder_ValidAskOrderWithNoResourcesGiven_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &pb.Order{
		//Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType:  pb.OrderType_ASK,
		SupplierID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:      "100",
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
		Id:         expectedID,
		SupplierID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:      "100",
		Slot: &pb.Slot{
			SupplierRating: 555,
			Resources:      &pb.Resources{},
		},
	}

	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:         expectedID,
			SupplierID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
			Price:      "100",
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

	// act
	obtained, err := m.CreateOrder(context.Background(), req)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMarketplaceCreateOrder_InValidRequest_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &pb.Order{
		Id:        "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		OrderType: pb.OrderType_BID,
		ByuerID:   "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:     "100",
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	expectedErr := fmt.Errorf("an error occured")

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

	// act
	_, err := m.CreateOrder(context.Background(), req)

	// assert
	assert.Error(t, err)
}

func TestMarketplaceCreateOrder_InvalidOrderTypeGiven_ErrorReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &pb.Order{
		Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:   "100",
	}

	m := NewMarketplace(nil, nil, nil)

	// act
	_, err := m.CreateOrder(context.Background(), req)

	// assert
	assert.EqualError(t, err, "incorrect order type given: ANY")
}
