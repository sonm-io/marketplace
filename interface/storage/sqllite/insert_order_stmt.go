package sqllite

import "github.com/gocraft/dbr"

func InsertOrderStmt(row interface{}) *dbr.InsertStmt {
	return dbr.InsertInto("orders").
		Columns("id",
			"type", "supplier_id", "buyer_id", "price",
			"slot_duration", "slot_buyer_rating", "slot_supplier_rating",
			"resources_cpu_cores", "resources_ram_bytes", "resources_gpu_count", "resources_storage",
			"resources_net_inbound", "resources_net_outbound", "resources_net_type", "resources_properties",
			"status").
		Record(row)
}
