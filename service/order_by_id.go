package service

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"

	"github.com/sonm-io/marketplace/mapper"
	mds "github.com/sonm-io/marketplace/mapper/datastruct"
)

// OrderByID retrieves an Order by the given ID.
func (ms *MarketService) OrderByID(ID string, result interface{}) error {
	res, ok := result.(*pb.Order)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	sql, args, err := ToSQL(OrderByIDStmt(ID))
	if err != nil {
		return err
	}

	var row mds.OrderRow
	if err := ms.s.FetchRow(&row, sql, args...); err != nil {
		if err == dbr.ErrNotFound {
			return fmt.Errorf("order %s is not found", ID)
		}
		return fmt.Errorf("cannot retrieve order: %v", err)
	}

	if row.Status != uint8(mds.Active) {
		return fmt.Errorf("order %s is inactive", ID)
	}

	order := ds.Order{Order: res}
	mapper.OrderFromRow(&order, &row)

	return nil
}
