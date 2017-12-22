package datastruct

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// OrderRows order row set.
type OrderRows []OrderRow

// OrderRow an order row.
type OrderRow struct {
	ID    string `db:"id"`
	Type  int32  `db:"type"`
	Price string `db:"price"`

	BuyerID    string `db:"buyer_id"`
	SupplierID string `db:"supplier_id"`

	Duration       uint64 `db:"slot_duration"`
	BuyerRating    int64  `db:"slot_buyer_rating"`
	SupplierRating int64  `db:"slot_supplier_rating"`

	CPUCores uint64 `db:"resources_cpu_cores"`
	RAMBytes uint64 `db:"resources_ram_bytes"`
	GPUCount uint64 `db:"resources_gpu_count"`
	Storage  uint64 `db:"resources_storage"`

	NetType     uint64 `db:"resources_net_type"`
	NetInbound  uint64 `db:"resources_net_inbound"`
	NetOutbound uint64 `db:"resources_net_outbound"`

	Properties Properties `db:"resources_properties"`

	Status uint8 `db:"status"`
}

// Properties represents Slot properties.
type Properties map[string]float64

// Value implements Valuer interface for database/sql.
func (p Properties) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan implements Scanner interface for database/sql.
func (p *Properties) Scan(src interface{}) error {
	if src == nil {
		*p = Properties{}
		return nil
	}

	value, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("resources_properties field must be a json-encoded []byte, got %T instead", src)
	}

	return json.Unmarshal(value, p)
}
