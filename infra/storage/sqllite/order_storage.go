package sqllite

import (
	"database/sql"
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	ds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"
)

// OrderStorage stores and retrieves Orders.
type OrderStorage struct {
	db     *sql.DB
	dbSess *dbr.Session
}

// NewOrderStorage creates an new instance of OrderStorage.
func NewOrderStorage(db *sql.DB) *OrderStorage {
	conn := &dbr.Connection{DB: db, EventReceiver: &dbr.NullEventReceiver{}, Dialect: dialect.SQLite3}
	dbSession := conn.NewSession(nil)

	return &OrderStorage{
		db:     db,
		dbSess: dbSession,
	}
}

// InsertRow inserts a row into orders table.
func (s *OrderStorage) InsertRow(row *ds.OrderRow) error {
	q := `
		INSERT OR REPLACE INTO orders
		(id, type, supplier_id, buyer_id, price, slot_duration, slot_buyer_rating, slot_supplier_rating,
		resources_cpu_cores, resources_ram_bytes, resources_gpu_count, resources_storage,
		resources_net_inbound, resources_net_outbound, resources_net_type, resources_properties)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.db.Exec(q,
		row.ID, row.Type, row.SupplierID, row.BuyerID, row.Price,
		row.Duration, row.BuyerRating, row.SupplierRating,
		row.CPUCores, row.RAMBytes, row.GPUCount, row.Storage,
		row.NetInbound, row.NetOutbound, row.NetType,
		row.Properties)
	if err != nil {
		return fmt.Errorf("cannot insert row: %v", err)
	}

	return nil
}

// DeleteRow removes a row with the given ID from orders table.
func (s *OrderStorage) DeleteRow(ID string) error {
	_, err := s.db.Exec("DELETE FROM orders WHERE id = ?", ID)
	if err != nil {
		return fmt.Errorf("cannot remove row: %v", err)
	}
	return nil
}

// UpdateStatus updates status field.
func (s *OrderStorage) UpdateStatus(ID string, status uint8) error {
	_, err := s.db.Exec("UPDATE orders SET status = ? WHERE id = ?", status, ID)
	if err != nil {
		return fmt.Errorf("cannot update status: %v", err)
	}
	return nil
}

// FetchRow fetches data by the given ID and maps it into row.
func (s *OrderStorage) FetchRow(row interface{}, query string, value ...interface{}) error {
	err := s.dbSess.SelectBySql(query, value...).LoadValue(row)
	return err
}

// FetchRows runs the given query and maps its result into rows.
func (s *OrderStorage) FetchRows(rows interface{}, query string, value ...interface{}) error {
	_, err := s.dbSess.SelectBySql(query, value...).Load(rows)
	return err
}
