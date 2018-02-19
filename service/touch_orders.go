package service

import "fmt"

// TouchOrders updates orders' TTL.
func (ms *MarketService) TouchOrders(IDs []string) error {
	stmt := TouchOrdersStmt(IDs)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := ms.s.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot update orders' ttl %v", err)
	}
	return nil
}
