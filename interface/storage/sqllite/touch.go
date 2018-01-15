package sqllite

import (
	"fmt"

	"github.com/gocraft/dbr"
)

// Touch updates given orders' TTL
func (s *OrderStorage) Touch(IDs []string) error {
	stmt := touchOrdersStmt(IDs)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := s.e.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot update orders' ttl %v", err)
	}
	return nil
}

func touchOrdersStmt(IDs []string) *dbr.UpdateStmt {
	return dbr.Update("orders").
		Set("status", Created).
		Where("id IN ?", IDs).
		Where("status != ?", Cancelled)
}
