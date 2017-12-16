package srv

import (
	"fmt"

	"github.com/pborman/uuid"
	"golang.org/x/net/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

// IDGenerator generates command IDs.
// Function is used to ease mocking.
var IDGenerator = func() string {
	return uuid.New()
}

// CreateOrder creates a bid order.
func (m *Marketplace) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	var (
		cmd intf.Command
		ID  string
	)
	switch req.OrderType {
	case pb.OrderType_BID:
		c := command.CreateBidOrder{}
		bindCreateBidOrderCommand(req, &c)
		ID = c.ID
		cmd = c

	case pb.OrderType_ASK:
		c := command.CreateAskOrder{}
		bindCreateAskOrderCommand(req, &c)
		ID = c.ID
		cmd = c

	default:
		return nil, fmt.Errorf("incorrect order type given: %v", req.OrderType.String())
	}

	if err := m.createOrder(ctx, cmd); err != nil {
		return nil, err
	}

	return m.GetOrderByID(ctx, &pb.ID{Id: ID})
}

func (m *Marketplace) createOrder(ctx context.Context, cmd intf.Command) error {
	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Creating order %+v", cmd)

	if err := m.commandBus.Handle(cmd); err != nil {
		logger.Sugar().Infof("cannot create order: %v\n", err)
		return fmt.Errorf("cannot create order: %v", err)
	}

	logger.Sugar().Info("order created")
	return nil
}

func bindCreateBidOrderCommand(req *pb.Order, cmd *command.CreateBidOrder) {
	// generate a unique ID if it's empty
	ID := req.GetId()
	if ID == "" {
		ID = IDGenerator()
	}

	cmd.ID = ID
	cmd.Price = req.GetPrice()
	cmd.BuyerID = req.GetByuerID()

	bindSlot(req.Slot, &cmd.Slot)
}

func bindCreateAskOrderCommand(req *pb.Order, cmd *command.CreateAskOrder) {
	// generate a unique ID if it's empty
	ID := req.GetId()
	if ID == "" {
		ID = IDGenerator()
	}
	cmd.ID = ID
	cmd.Price = req.GetPrice()
	cmd.SupplierID = req.GetByuerID()

	bindSlot(req.Slot, &cmd.Slot)
}

func bindSlot(pbSlot *pb.Slot, dsSlot *ds.Slot) {
	if pbSlot == nil {
		return
	}

	dsSlot.SupplierRating = pbSlot.GetSupplierRating()
	dsSlot.BuyerRating = pbSlot.GetBuyerRating()

	if pbSlot.Resources == nil {
		return
	}

	res := pbSlot.GetResources()
	dsSlot.Resources = ds.Resources{
		CPUCores:      res.GetCpuCores(),
		RAMBytes:      res.GetRamBytes(),
		GPUCount:      ds.GPUCount(res.GetGpuCount()),
		Storage:       res.GetStorage(),
		NetworkType:   ds.NetworkType(res.GetNetworkType()),
		NetTrafficIn:  res.GetNetTrafficIn(),
		NetTrafficOut: res.GetNetTrafficOut(),
		Properties:    res.GetProperties(),
	}
}
