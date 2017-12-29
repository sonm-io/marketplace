package sqllite

import "github.com/gocraft/dbr"

// GetOrderByIDStmt is a factory method that builds a statement to get order by ID.
func GetOrderByIDStmt(ID string) (*dbr.SelectStmt, error) {
	stmt := dbr.Select("id", "type", "supplier_id", "buyer_id", "price",
		"slot_duration", "slot_buyer_rating", "slot_supplier_rating",
		"resources_cpu_cores", "resources_ram_bytes", "resources_gpu_count", "resources_storage",
		"resources_net_inbound", "resources_net_outbound", "resources_net_type", "resources_properties",
		"status").
		From("orders").
		Where("id = ?", ID)

	return stmt, nil
}
