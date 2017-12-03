package srv

import (
	"fmt"

	"github.com/pborman/uuid"
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

// CreateOrder creates a bid order.
func (m *Marketplace) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	var cmd command.CreateBidOrder
	if err := m.bind(req, &cmd); err != nil {
		return nil, fmt.Errorf("cannot map request to command: %v", err)
	}

	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Creating bid order %+v", cmd)

	if err := m.commandBus.Handle(cmd); err != nil {
		logger.Sugar().Infof("cannot create bid order: %v\n", err)
		return nil, fmt.Errorf("cannot create bid order: %v", err)
	}

	logger.Sugar().Infof("bid order %s created\n", cmd.ID)

	return m.GetOrderByID(ctx, &pb.ID{Id: cmd.ID})
}

func (m *Marketplace) bind(req *pb.Order, cmd *command.CreateBidOrder) error {

	// get id from request or generate new
	ID := req.GetId()
	if ID == "" {
		ID = uuid.New()
	}

	c := &command.CreateBidOrder{
		ID:      ID,
		Price:   req.GetPrice(),
		BuyerID: req.GetByuerID(),
	}

	*cmd = *c

	if req.Slot != nil {
		cmd.Slot.SupplierRating = req.GetSlot().GetSupplierRating()
		cmd.Slot.BuyerRating = req.GetSlot().GetBuyerRating()
		if req.Slot.Resources != nil {
			res := req.GetSlot().GetResources()
			cmd.Slot.Resources = ds.Resources{
				CPUCores: res.GetCpuCores(),
				RAMBytes: res.GetRamBytes(),
				Storage:  res.GetStorage(),
			}
		}
	}

	return nil
}
