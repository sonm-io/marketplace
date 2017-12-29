package sqllite

import "github.com/gocraft/dbr"

func CancelOrderStmt(ID string) *dbr.UpdateStmt {
	return dbr.Update("orders").
		Set("status", InActive).
		Where("id = ?", ID)
}
