package service

import "fmt"

// CancelOrder marks the given order as cancelled.
func (ms *MarketService) CancelOrder(ID string) error {
	query, args, _ := ToSQL(CancelOrderStmt(ID))

	if err := ms.s.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot cancel order: %v", err)
	}
	return nil
}
