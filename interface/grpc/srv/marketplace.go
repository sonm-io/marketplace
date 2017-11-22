package srv

import (
	"github.com/pborman/uuid"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/command"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"

	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
	"golang.org/x/net/context"
)

//const (
//defaultResultsCount = 100
//)

// read side
type OrderByID interface {
	Handle(req intf.Query, result interface{}) error
}

type Marketplace struct {
	commandBus intf.CommandHandler
	orderByID  OrderByID
}

func NewMarketplace(c intf.CommandHandler, q OrderByID) *Marketplace {
	return &Marketplace{
		commandBus: c,
		orderByID:  q,
	}
}

func (m *Marketplace) GetOrderByID(_ context.Context, req *pb.ID) (*pb.Order, error) {
	order := &query.GetOrderResult{}
	if err := m.orderByID.Handle(query.GetOrder{ID:req.Id}, order); err != nil {
		return nil, err
	}

	// build result
	resp := &pb.Order{
		Id:         order.ID,
		ByuerID:    order.BuyerID,
		SupplierID: order.SupplierID,
		Price:      order.Price,
		//OrderType:
		//Slot:
	}

	return resp, nil
}

func (m *Marketplace) CancelOrder(_ context.Context, req *pb.Order) (*pb.Empty, error) {
	if err := m.commandBus.Handle(command.CancelOrder{ID: req.Id}); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (m *Marketplace) CreateOrder(_ context.Context, req *pb.Order) (*pb.Order, error) {

	// map request to command
	// TODO: (screwyprof) move to smth like cmd := model.Bind(req), or model.bind(&req, &cmd)
	 cmd := command.CreateOrder{
		ID:         uuid.New(),
		SupplierID: req.SupplierID,
		BuyerID:    req.ByuerID,
		Price:      req.Price,
	}

	// handle command
	if err := m.commandBus.Handle(cmd); err != nil {
		return nil, err
	}

	// return response
	return &pb.Order{Id: cmd.ID}, nil
}

func (m *Marketplace) GetOrders(_ context.Context, req *pb.GetOrdersRequest) (*pb.GetOrdersReply, error) {
	//slot, err := structs.NewSlot(req.Slot)
	//if err != nil {
	//	return nil, err
	//}
	//
	//resultCount := req.GetCount()
	//if resultCount == 0 {
	//	resultCount = defaultResultsCount
	//}
	//
	//searchParams := &searchParams{
	//	slot:      slot,
	//	orderType: req.GetOrderType(),
	//	count:     resultCount,
	//}
	//
	//orders, err := m.s.GetOrders(searchParams)
	//if err != nil {
	//	return nil, err
	//}
	//
	//innerOrders := []*pb.Order{}
	//for _, o := range orders {
	//	innerOrders = append(innerOrders, o.Unwrap())
	//}
	//
	//return &pb.GetOrdersReply{
	//	Orders: innerOrders,
	//}, nil
	return &pb.GetOrdersReply{}, nil
}

func (m *Marketplace) GetProcessing(ctx context.Context, req *pb.Empty) (*pb.GetProcessingReply, error) {
	// This method exists just to match the Marketplace interface.
	// The Market service itself is unable to know anything about processing orders.
	// This method is implemented for Node in `insonmnia/node/market.go:348`
	return nil, nil
}
