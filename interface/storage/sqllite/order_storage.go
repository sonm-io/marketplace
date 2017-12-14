package sqllite

import (
	"database/sql"
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
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

// Add adds the given Order to the storage.
func (s *OrderStorage) Add(o *ds.Order) error {
	if o.Slot == nil {
		o.Slot = &ds.Slot{}
	}

	q := `
		INSERT OR REPLACE INTO orders
		(id, type, supplier_id, buyer_id, price, slot_buyer_rating, slot_supplier_rating,
		resources_cpu_cores, resources_ram_bytes, resources_gpu_count, resources_storage,
		resources_net_inbound, resources_net_outbound, resources_net_type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.e.Exec(q,
		o.ID, int(o.OrderType), o.SupplierID, o.BuyerID, o.Price, o.Slot.BuyerRating, o.Slot.SupplierRating,
		o.Slot.Resources.CPUCores, o.Slot.Resources.RAMBytes, int(o.Slot.Resources.GPUCount), o.Slot.Resources.Storage,
		o.Slot.Resources.NetTrafficIn, o.Slot.Resources.NetTrafficOut, int(o.Slot.Resources.NetworkType))
	if err != nil {
		return fmt.Errorf("cannot add new order: %v", err)
	}

	return nil
}

// Remove removes an Order with the given ID from OrderStorage.
// If no orders found, an error is returned.
func (s *OrderStorage) Remove(ID string) error {
	_, err := s.e.Exec("DELETE FROM orders WHERE id = ?", ID)
	if err != nil {
		return fmt.Errorf("cannot remove order: %v", err)
	}
	return nil
}

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderStorage) ByID(ID string) (ds.Order, error) {

	var row OrderRow

	err := s.e.QueryRow(
		`SELECT id, type, supplier_id, buyer_id, price, slot_buyer_rating, slot_supplier_rating,
			   		resources_cpu_cores, resources_ram_bytes, resources_gpu_count, resources_storage,
			   		resources_net_inbound, resources_net_outbound, resources_net_type
			   FROM orders
			   WHERE id = ?`, ID).
		Scan(&row.ID, &row.Type, &row.SupplierID, &row.BuyerID, &row.Price, &row.BuyerRating, &row.SupplierRating,
			&row.CPUCores, &row.RAMBytes, &row.GPUCount, &row.Storage,
			&row.NetInbound, &row.NetOutbound, &row.NetType)

	if err != nil {
		return ds.Order{}, fmt.Errorf("cannot get order: %v", err)
	}

	// a kludge
	row.Properties = map[string]float64{
		"hash_rate": 105.7,
	}

	order := ds.Order{}
	mapOrder(&order, &row)

	return order, nil
}

func mapOrder(order *ds.Order, row *OrderRow) {
	if order == nil {
		return
	}

	order.ID = row.ID
	order.OrderType = ds.OrderType(row.Type)
	order.Price = row.Price

	order.BuyerID = row.BuyerID
	order.SupplierID = row.SupplierID

	if order.Slot == nil {
		order.Slot = &ds.Slot{}
	}

	slot := &ds.Slot{
		BuyerRating:    row.BuyerRating,
		SupplierRating: row.SupplierRating,

		Resources: ds.Resources{
			CPUCores: row.CPUCores,
			RAMBytes: row.RAMBytes,
			GPUCount: ds.GPUCount(row.GPUCount),
			Storage:  row.Storage,

			NetworkType:   ds.NetworkType(row.NetType),
			NetTrafficIn:  row.NetInbound,
			NetTrafficOut: row.NetOutbound,

			Properties: row.Properties,
		},
	}

	order.Slot = slot
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) ([]ds.Order, error) {

	//rows, err := db.Query("select id, name from users where id = ?", 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err := rows.Scan(&id, &name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(id, name)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//b := inmemory.NewBuilder()
	//b.WithLimit(limit)
	//b.WithSpec(spec)
	//
	//elements, err := s.e.Match(b.Build())
	//if err != nil {
	//	return nil, err
	//}
	//
	//var orders []ds.Order
	//for _, el := range elements {
	//	order := el.(*ds.Order)
	//	orders = append(orders, *order)
	//}
	//
	//sort.Sort(ByPrice(orders))
	//
	//return orders, nil
	return nil, nil
}
