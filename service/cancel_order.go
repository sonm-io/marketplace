package service

import "fmt"

// CancelOrder marks the given order as cancelled.
func (ms *MarketService) CancelOrder(ID string) error {
	stmt := CancelOrderStmt(ID)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := ms.s.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot cancel order: %v", err)
	}
	return nil
}
