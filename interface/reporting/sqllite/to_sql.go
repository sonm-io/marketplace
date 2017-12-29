package sqllite

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

func ToSQL(stmt dbr.Builder) (string, []interface{}, error) {
	if stmt == nil {
		return "", nil, fmt.Errorf("cannot build sql: stmt is nil")
	}

	buf := dbr.NewBuffer()
	if err := stmt.Build(dialect.SQLite3, buf); err != nil {
		return "", nil, fmt.Errorf("cannot build sql statement: %v", err)
	}
	return buf.String(), buf.Value(), nil
}
