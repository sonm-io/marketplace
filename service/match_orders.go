package service

import (
	"fmt"

	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/proto"

	"github.com/sonm-io/marketplace/mapper"
	mds "github.com/sonm-io/marketplace/mapper/datastruct"
)

// MatchOrders
// Retrieves Orders by the given Spec.
//
// if req.Limit is > 0, then only the given number of Orders will be returned.
func (ms *MarketService) MatchOrders(req *pb.Order, limit uint64, result interface{}) error {
	res, ok := result.(*pb.GetOrdersReply)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	stmt, err := MatchOrdersStmt(ds.Order{Order: req}, limit)
	if err != nil {
		return err
	}

	sql, args, _ := ToSQL(stmt)

	var rows mds.OrderRows
	if err := ms.s.FetchRows(&rows, sql, args...); err != nil {
		return fmt.Errorf("cannot match orders: %v", err)
	}

	ms.filterOrders(req, rows, res)

	return nil
}

func (ms *MarketService) filterOrders(req *pb.Order, rows mds.OrderRows, res *pb.GetOrdersReply) {
	var order ds.Order
	propertiesSpec := NewPropertiesSpec(req)
	for idx := range rows {
		order = ds.Order{Order: &pb.Order{}}
		mapper.OrderFromRow(&order, &rows[idx])

		if !propertiesSpec.IsSatisfiedBy(order) {
			continue
		}
		(*res).Orders = append((*res).Orders, order.Order)
	}
}
