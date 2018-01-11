package srv

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/intf/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestMarketplaceGetOrdersByID_BuyerIDGiven_CorrespondentOrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			ByuerID: buyerID,
		},
		Count: 100,
	}

	q := query.GetOrders{}
	bindGetOrdersQuery(req, &q)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
				ByuerID: buyerID,
				Price:   "100",
				Slot: &pb.Slot{
					Duration: 900,
					Resources: &pb.Resources{
						CpuCores: 4,
						RamBytes: 10000,
					},
				},
			},
		},
	}

	ordersReport := []report.GetOrderReport{
		{
			Order: ds.Order{
				ID:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
				BuyerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
				Price:   "100",
				Slot: &ds.Slot{
					Duration: 900,
					Resources: ds.Resources{
						CPUCores: 4,
						RAMBytes: 10000,
					},
				},
			},
		},
	}

	ordersBySpec := mocks.NewMockQueryHandler(ctrl)
	ordersBySpec.EXPECT().Handle(q, &report.GetOrdersReport{}).
		SetArg(1, ordersReport).
		Return(nil)

	m := NewMarketplace(nil, nil, ordersBySpec)

	obtained, err := m.GetOrders(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}

func TestMarketplaceGetOrdersByID_SlotGiven_CorrespondentOrdersReturned(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buyerID := "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db"
	req := &pb.GetOrdersRequest{
		Order: &pb.Order{
			OrderType: pb.OrderType_BID,
			ByuerID:   buyerID,
			Slot: &pb.Slot{
				Resources: &pb.Resources{
					CpuCores: 2,
				},
			},
		},
		Count: 100,
	}

	q := query.GetOrders{}
	bindGetOrdersQuery(req, &q)

	expected := &pb.GetOrdersReply{
		Orders: []*pb.Order{
			{
				Id:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
				ByuerID: buyerID,
				Price:   "100",
				Slot: &pb.Slot{
					Duration: 900,
					Resources: &pb.Resources{
						CpuCores: 4,
						RamBytes: 10000,
					},
				},
			},
		},
	}

	ordersReport := []report.GetOrderReport{
		{
			Order: ds.Order{
				ID:      "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
				BuyerID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
				Price:   "100",
				Slot: &ds.Slot{
					Duration: 900,
					Resources: ds.Resources{
						CPUCores: 4,
						RAMBytes: 10000,
					},
				},
			},
		},
	}

	ordersBySpec := mocks.NewMockQueryHandler(ctrl)
	ordersBySpec.EXPECT().Handle(q, &report.GetOrdersReport{}).
		SetArg(1, ordersReport).
		Return(nil)

	m := NewMarketplace(nil, nil, ordersBySpec)

	obtained, err := m.GetOrders(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}
