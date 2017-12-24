package inmemory

import (
	"github.com/sonm-io/marketplace/infra/storage/inmemory"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
	"sort"
)

// OrderStorage stores and retrieves Orders.
type OrderStorage struct {
	e Engine
}

// NewOrderStorage creates an new instance of OrderStorage.
func NewOrderStorage(e Engine) *OrderStorage {
	return &OrderStorage{
		e: e,
	}
}

// Add adds the given Order to the storage.
func (s *OrderStorage) Add(o *ds.Order) error {
	return s.e.Add(o, o.ID)
}

// Remove removes an Order with the given ID from OrderStorage.
// If no orders found, an error is returned.
func (s *OrderStorage) Remove(ID string) error {
	return s.e.Remove(ID)
}

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderStorage) ByID(ID string) (ds.Order, error) {

	el, err := s.e.Get(ID)
	if err != nil {
		return ds.Order{}, err
	}
	order := el.(*ds.Order)
	return *order, nil
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) ([]ds.Order, error) {
	// spec is empty, nothing to return
	if spec == nil {
		return nil, nil
	}

	b := inmemory.NewBuilder()
	b.WithLimit(limit)
	b.WithSpec(spec)

	elements, err := s.e.Match(b.Build())
	if err != nil {
		return nil, err
	}

	var orders []ds.Order
	for _, el := range elements {
		order := el.(*ds.Order)
		orders = append(orders, *order)
	}

	sort.Sort(ByPrice(orders))

	return orders, nil
}

// ByPrice implements sort.Interface; it allows for sorting Orders by Price field.
type ByPrice []ds.Order

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }
