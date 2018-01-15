package sqllite

import (
	"fmt"

	"github.com/gocraft/dbr"
)

// Cancel marks an Order with the given ID as InActive.
func (s *OrderStorage) Cancel(ID string) error {
	stmt := cancelOrderStmt(ID)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := s.e.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot remove order: %v", err)
	}
	return nil
}

func cancelOrderStmt(ID string) *dbr.UpdateStmt {
	return dbr.Update("orders").
		Set("status", Cancelled).
		Where("id = ?", ID)
}
