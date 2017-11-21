package inmemory

import (
	"errors"
	"sync"

	"github.com/sonm-io/marketplace/entity"
)

var (
	errOrderNotFound     = errors.New("order cannot be found")
//	errPriceIsZero       = errors.New("order price cannot be less or equal than zero")
//	errOrderIsNil        = errors.New("order cannot be nil")
//	errSlotIsNil         = errors.New("order slot cannot be nil")
//	errResourcesIsNil    = errors.New("slot resources cannot be nil")
//	errSearchParamsIsNil = errors.New("search params cannot be nil")
)

type OrderStorage struct {
	sync.RWMutex
	db map[string]*entity.Order
}

func NewStorage() *OrderStorage {
	return &OrderStorage{
		db: make(map[string]*entity.Order),
	}
}

func (s *OrderStorage) ByID(id string) (*entity.Order, error) {
	s.RLock()
	defer s.RUnlock()

	ord, ok := s.db[id]
	if !ok {
		return nil, errOrderNotFound
	}

	return ord, nil
}

func (s *OrderStorage) Store(o *entity.Order) error {
	s.Lock()
	defer s.Unlock()

	s.db[o.ID] = o
	return nil
}

func (s *OrderStorage) Remove(ID string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.db[ID]; !ok {
		return errOrderNotFound
	}

	delete(s.db, ID)
	return nil
}

// searchParams holds all fields that are used to search on the market.
// Extend this structure instead of increasing amount of params accepted
// by OrderStorage.GetOrders() function.
//type searchParams struct {
//	slot      *entity.Slot
//	orderType pb.OrderType
//	count     uint64
//}

//func (s *OrderStorage) generateID() string {
//	return uuid.New()
//}

//
//func (s *OrderStorage) GetOrders(c *searchParams) ([]*entity.Order, error) {
//	if c == nil {
//		return nil, errSearchParamsIsNil
//	}
//
//	if c.slot == nil {
//		return nil, errSlotIsNil
//	}
//
//	s.RLock()
//	defer s.RUnlock()
//
//	orders := []*entity.Order{}
//	for _, order := range s.db {
//		if uint64(len(orders)) >= c.count {
//			break
//		}
//
//		if compareOrderAndSlot(c.slot, order, c.orderType) {
//			orders = append(orders, order)
//		}
//	}
//
//	sort.Sort(entity.ByPrice(orders))
//	return orders, nil
//}

//func compareOrderAndSlot(slot *entity.Slot, order *entity.Order, typ pb.OrderType) bool {
//	if typ != pb.OrderType_ANY && typ != order.GetType() {
//		return false
//	}
//
//	os, _ := entity.NewSlot(order.Unwrap().Slot)
//	return slot.Compare(os, order.GetType())
//}
