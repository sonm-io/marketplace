package srv

import (
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ds "github.com/sonm-io/marketplace/datastruct"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
)

// CreateOrder creates a bid order.
func (m *Marketplace) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	if err := CheckPermissions(ctx, req); err != nil {
		return nil, err
	}

	if err := m.validate(req); err != nil {
		return nil, err
	}

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
		return nil, status.Errorf(codes.Internal, "Cannot create order: %v", err)
	}

	return m.GetOrderByID(ctx, &pb.ID{Id: ID})
}

func (m *Marketplace) validate(req *pb.Order) error {
	if req == nil || req.Slot == nil || req.Slot.Resources == nil {
		return nil
	}

	if req.Slot.Resources.GpuCount == pb.GPUCount_SINGLE_GPU {
		return fmt.Errorf("SINGLE_GPU has been deprecated, only NO_GPU and MULTIPLE_GPU are allowed")
	}

	return nil
}

func (m *Marketplace) createOrder(ctx context.Context, cmd intf.Command) error {
	logger := ctx_zap.Extract(ctx)
	logger.Sugar().Infof("Creating order %+v", cmd)

	if err := m.commandBus.Handle(cmd); err != nil {
		return err
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
	cmd.BuyerID = req.GetByuerID()
	cmd.SupplierID = req.GetSupplierID()

	reqPricePerSecond := req.GetPricePerSecond()
	if reqPricePerSecond != nil {
		cmd.Price = reqPricePerSecond.Unwrap().String()
	}

	bindSlot(req.Slot, &cmd.Slot)
}

func bindCreateAskOrderCommand(req *pb.Order, cmd *command.CreateAskOrder) {
	// generate a unique ID if it's empty
	ID := req.GetId()
	if ID == "" {
		ID = IDGenerator()
	}
	cmd.ID = ID
	cmd.SupplierID = req.GetSupplierID()
	cmd.BuyerID = req.GetByuerID()

	reqPricePerSecond := req.GetPricePerSecond()
	if reqPricePerSecond != nil {
		cmd.Price = reqPricePerSecond.Unwrap().String()
	}

	bindSlot(req.Slot, &cmd.Slot)
}

func bindSlot(pbSlot *pb.Slot, dsSlot *ds.Slot) {
	if pbSlot == nil {
		return
	}

	dsSlot.SupplierRating = pbSlot.GetSupplierRating()
	dsSlot.BuyerRating = pbSlot.GetBuyerRating()
	dsSlot.Duration = pbSlot.GetDuration()

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
