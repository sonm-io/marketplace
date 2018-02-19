package service

import (
	"fmt"
	"github.com/sonm-io/marketplace/ds"
	"github.com/sonm-io/marketplace/mapper"
	mds "github.com/sonm-io/marketplace/mapper/datastruct"
)

func (ms *MarketService) createOrder(order *ds.Order) error {
	row := mds.OrderRow{}
	mapper.OrderToRow(order, &row)
	row.Status = mds.Active

	stmt := InsertOrderStmt(row)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := ms.s.InsertRow(query, args...); err != nil {
		return fmt.Errorf("cannot create a new order: %v", err)
	}
	return nil
}
