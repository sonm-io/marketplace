package sqllite

//import (
//	"encoding/json"
//	"fmt"
//)

// OrderRow an order row.
type OrderRow struct {
	ID    string `db:"id"`
	Type  int32  `db:"type"`
	Price int64  `db:"price"`

	BuyerID    string `db:"buyer_id"`
	SupplierID string `db:"supplier_id"`

	BuyerRating    int64 `db:"slot_buyer_rating"`
	SupplierRating int64 `db:"slot_supplier_rating"`

	CPUCores uint64 `db:"resources_cpu_cores"`
	RAMBytes uint64 `db:"resources_ram_bytes"`
	GPUCount uint64 `db:"resources_gpu_count"`
	Storage  uint64 `db:"resources_storage"`

	NetType     uint64 `db:"resources_net_type"`
	NetInbound  uint64 `db:"resources_net_inbound"`
	NetOutbound uint64 `db:"resources_net_outbound"`

	Properties Properties `db:"resources_properties" json:"properties"`
}

// Properties represents Slot properties.
type Properties map[string]float64
