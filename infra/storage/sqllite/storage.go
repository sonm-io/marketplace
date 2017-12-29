package sqllite

import (
	"database/sql"
	"strings"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

// Storage CRUD interface.
type Storage struct {
	dbSess *dbr.Session
}

// NewStorage creates an new instance of Storage.
func NewStorage(db *sql.DB) *Storage {
	conn := &dbr.Connection{DB: db, EventReceiver: &dbr.NullEventReceiver{}, Dialect: dialect.SQLite3}
	return &Storage{
		dbSess: conn.NewSession(nil),
	}
}

// InsertRow inserts a row into the table.
func (s *Storage) InsertRow(query string, args ...interface{}) error {
	// builder generates INSERT to, but INSERT OR REPLACE is NEEDED
	query = strings.Replace(query, "INSERT INTO", "INSERT OR REPLACE INTO", 1)
	_, err := s.dbSess.InsertBySql(query, args...).Exec()
	return err
}

// UpdateRow updates given row.
func (s *Storage) UpdateRow(query string, value ...interface{}) error {
	_, err := s.dbSess.UpdateBySql(query, value...).Exec()
	return err
}

// FetchRow fetches data by the given ID and maps it into row.
func (s *Storage) FetchRow(row interface{}, query string, value ...interface{}) error {
	err := s.dbSess.SelectBySql(query, value...).LoadValue(row)
	return err
}

// FetchRows runs the given query and maps its result into rows.
func (s *Storage) FetchRows(rows interface{}, query string, value ...interface{}) error {
	_, err := s.dbSess.SelectBySql(query, value...).Load(rows)
	return err
}
