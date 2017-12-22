package sqllite

import (
	"database/sql"
	"fmt"

	ds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"
)

// OrderStorage stores and retrieves Orders.
type OrderStorage struct {
	e *sql.DB
}

// NewOrderStorage creates an new instance of OrderStorage.
func NewOrderStorage(e *sql.DB) *OrderStorage {
	return &OrderStorage{
		e: e,
	}
}

// InsertRow inserts a row into orders table.
func (s *OrderStorage) InsertRow(row *ds.OrderRow) error {
	q := `
		INSERT OR REPLACE INTO orders
		(id, type, supplier_id, buyer_id, price, slot_buyer_rating, slot_supplier_rating,
		resources_cpu_cores, resources_ram_bytes, resources_gpu_count, resources_storage,
		resources_net_inbound, resources_net_outbound, resources_net_type, resources_properties)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.e.Exec(q,
		row.ID, row.Type, row.SupplierID, row.BuyerID, row.Price, row.BuyerRating, row.SupplierRating,
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
	_, err := s.e.Exec("DELETE FROM orders WHERE id = ?", ID)
	if err != nil {
		return fmt.Errorf("cannot remove row: %v", err)
	}
	return nil
}

// UpdateStatus updates status field.
func (s *OrderStorage) UpdateStatus(ID string, status uint8) error {
	_, err := s.e.Exec("UPDATE orders SET status = ? WHERE id = ?", status, ID)
	if err != nil {
		return fmt.Errorf("cannot update status: %v", err)
	}
	return nil
}

// FetchRow fetches data by the given ID and maps it into row.
func (s *OrderStorage) FetchRow(ID string, row *ds.OrderRow) error {
	err := s.e.QueryRow(
		`SELECT id, type, supplier_id, buyer_id, price, slot_buyer_rating, slot_supplier_rating,
			   		resources_cpu_cores, resources_ram_bytes, resources_gpu_count, resources_storage,
			   		resources_net_inbound, resources_net_outbound, resources_net_type, resources_properties,
			   		status
			   FROM orders
			   WHERE id = ?`, ID).
		Scan(&row.ID, &row.Type, &row.SupplierID, &row.BuyerID, &row.Price, &row.BuyerRating, &row.SupplierRating,
			&row.CPUCores, &row.RAMBytes, &row.GPUCount, &row.Storage,
			&row.NetInbound, &row.NetOutbound, &row.NetType, &row.Properties, &row.Status)

	return err
}

// FetchAll fetches all the rows.
func (s *OrderStorage) FetchAll() (ds.OrderRows, error) {
	rows, err := s.e.Query(
		`SELECT id, type, supplier_id, buyer_id, price, slot_buyer_rating, slot_supplier_rating,
			   		resources_cpu_cores, resources_ram_bytes, resources_gpu_count, resources_storage,
			   		resources_net_inbound, resources_net_outbound, resources_net_type, resources_properties
			   FROM orders
			   WHERE status = ?
			   ORDER BY price`, 1) // only active rows
	if err != nil {
		return nil, fmt.Errorf("cannot get orders: %v", err)
	}
	defer rows.Close()

	var (
		row       ds.OrderRow
		orderRows ds.OrderRows
	)

	for rows.Next() {
		row = ds.OrderRow{}

		err := rows.Scan(&row.ID, &row.Type, &row.SupplierID, &row.BuyerID, &row.Price,
			&row.BuyerRating, &row.SupplierRating,
			&row.CPUCores, &row.RAMBytes, &row.GPUCount, &row.Storage,
			&row.NetInbound, &row.NetOutbound, &row.NetType,
			&row.Properties)
		if err != nil {
			return nil, fmt.Errorf("cannot scan order raw into struct: %v", err)
		}

		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("cannot retrieve orders: %v", err)
		}
		orderRows = append(orderRows, row)
	}

	return orderRows, nil
}
