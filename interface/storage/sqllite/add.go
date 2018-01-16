package sqllite

import (
	"fmt"

	"github.com/gocraft/dbr"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/interface/mapper"
	sds "github.com/sonm-io/marketplace/interface/mapper/datastruct"
)

// Add adds the given Order to the storage.
func (s *OrderStorage) Add(o *ds.Order) error {
	if o == nil {
		return fmt.Errorf("cannot add an empty order")
	}

	row := sds.OrderRow{}
	mapper.OrderToRow(o, &row)
	row.Status = Created

	stmt := insertOrderStmt(row)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := s.e.InsertRow(query, args...); err != nil {
		return fmt.Errorf("cannot add new order: %v", err)
	}
	return nil
}

func insertOrderStmt(row interface{}) *dbr.InsertStmt {
	return dbr.InsertInto("orders").
		Columns("id",
			"type", "supplier_id", "buyer_id", "price",
			"slot_duration", "slot_buyer_rating", "slot_supplier_rating",
			"resources_cpu_cores", "resources_ram_bytes", "resources_gpu_count", "resources_storage",
			"resources_net_inbound", "resources_net_outbound", "resources_net_type", "resources_properties",
			"status").
		Record(row)
}
