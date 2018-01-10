package srv

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/intf/mocks"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

func TestMarketplace_GetOrderByID(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pricePerSecond, err := pb.NewBigIntFromString("100")
	require.NoError(t, err)

	req := &pb.ID{Id: "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7"}
	expected := &pb.Order{
		Id:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
		ByuerID:        "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
		PricePerSecond: pricePerSecond,
		Slot: &pb.Slot{
			Duration: 900,
			Resources: &pb.Resources{
				CpuCores: 4,
				RamBytes: 10000,
			},
		},
	}

	q := query.GetOrder{
		ID: req.GetId(),
	}

	orderReport := report.GetOrderReport{
		Order: ds.Order{
			ID:             "cfef34ae-58d3-4693-8c6c-d1b95e7ed7e7",
			BuyerID:        "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
			PricePerSecond: "100",
			Slot: &ds.Slot{
				Duration: 900,
				Resources: ds.Resources{
					CPUCores: 4,
					RAMBytes: 10000,
				},
			},
		},
	}

	orderByIDMock := mocks.NewMockQueryHandler(ctrl)
	orderByIDMock.EXPECT().Handle(q, &report.GetOrderReport{}).
		SetArg(1, orderReport).
		Return(nil)

	m := NewMarketplace(nil, orderByIDMock, nil)

	obtained, err := m.GetOrderByID(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expected, obtained)
}
