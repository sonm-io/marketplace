package srv

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"fmt"
	"github.com/sonm-io/marketplace/usecase/intf/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestMarketplaceCreateOrder_ValidRequest_ValidResponse(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &pb.Order{
		Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:   100,
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
		Price:   100,
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
			Price:   100,
			Slot: &ds.Slot{
				Resources: ds.Resources{
					CPUCores: 4,
					RAMBytes: 10000,
				},
			},
		},
	}

	var cmd command.CreateBidOrder
	bindCreateOrderCommand(req, &cmd)

	createOrderMock := mocks.NewMockCommandHandler(ctrl)
	createOrderMock.EXPECT().Handle(cmd).
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
		Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		Price:   100,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	expectedErr := fmt.Errorf("an error occured")

	var cmd command.CreateBidOrder
	bindCreateOrderCommand(req, &cmd)

	createOrderMock := mocks.NewMockCommandHandler(ctrl)
	createOrderMock.EXPECT().Handle(cmd).
		Return(expectedErr)

	m := NewMarketplace(createOrderMock, nil, nil)

	// act
	_, err := m.CreateOrder(context.Background(), req)

	// assert
	assert.Error(t, err)
}
