package service

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"

	"github.com/sonm-io/marketplace/ds"
	mds "github.com/sonm-io/marketplace/mapper/datastruct"
	"github.com/sonm-io/marketplace/service/filter"
)

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

func CancelOrderStmt(ID string) *dbr.UpdateStmt {
	return dbr.Update("orders").
		Set("status", mds.Cancelled).
		Where("id = ?", ID)
}

func TouchOrdersStmt(IDs []string) *dbr.UpdateStmt {
	return dbr.Update("orders").
		Set("status", mds.Active).
		Where("id IN ?", IDs).
		Where("status != ?", mds.Cancelled)
}

func OrderByIDStmt(ID string) *dbr.SelectStmt {
	stmt := dbr.Select("id", "type", "supplier_id", "buyer_id", "price",
		"slot_duration", "slot_buyer_rating", "slot_supplier_rating",
		"resources_cpu_cores", "resources_ram_bytes", "resources_gpu_count", "resources_storage",
		"resources_net_inbound", "resources_net_outbound", "resources_net_type", "resources_properties",
		"status").
		From("orders").
		Where("id = ?", ID).
		Where("status != ?", mds.Expired)

	return stmt
}

func MatchOrdersStmt(order ds.Order, limit uint64) (*dbr.SelectStmt, error) {
	cond, err := filter.MatchOrder(order)
	if err != nil {
		return nil, fmt.Errorf("cannot build conditions: %v", err)
	}

	// select active and unexpired records which satisfy the condition.
	stmt := dbr.Select("id", "type", "supplier_id", "buyer_id", "price",
		"slot_duration", "slot_buyer_rating", "slot_supplier_rating",
		"resources_cpu_cores", "resources_ram_bytes", "resources_gpu_count", "resources_storage",
		"resources_net_inbound", "resources_net_outbound", "resources_net_type", "resources_properties").
		From("orders").
		Where("status = ?", mds.Active).
		Where(cond).
		OrderAsc("price")

	if limit > 0 {
		stmt.Limit(limit)
	}

	return stmt, nil
}

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
