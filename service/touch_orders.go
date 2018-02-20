package service

import "fmt"

// TouchOrders updates orders' TTL.
func (ms *MarketService) TouchOrders(IDs []string) error {
	query, args, _ := ToSQL(TouchOrdersStmt(IDs))

	if err := ms.s.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot update orders' ttl %v", err)
	}
	return nil
}
