package sqllite

import (
	"fmt"

	"github.com/gocraft/dbr"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/interface/reporting/sqllite/filter"
)

// MatchOrdersStmt is a factory method that builds a statement to get orders by filter.
func MatchOrdersStmt(order ds.Order, limit uint64) (*dbr.SelectStmt, error) {
	cond, err := filter.MatchOrder(order)
	if err != nil {
		return nil, fmt.Errorf("cannot build conditions: %v", err)
	}

	stmt := dbr.Select("id", "type", "supplier_id", "buyer_id", "price",
		"slot_duration", "slot_buyer_rating", "slot_supplier_rating",
		"resources_cpu_cores", "resources_ram_bytes", "resources_gpu_count", "resources_storage",
		"resources_net_inbound", "resources_net_outbound", "resources_net_type", "resources_properties").
		From("orders").
		Where("status = ?", Active).
		Where(cond).
		OrderAsc("price")

	if limit > 0 {
		stmt.Limit(limit)
	}

	return stmt, nil
}
