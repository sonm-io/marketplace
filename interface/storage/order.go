package storage

import (
	"github.com/sonm-io/marketplace/entity"
)

//var (
//	errOrderNotFound     = errors.New("order cannot be found")
//	errPriceIsZero       = errors.New("order price cannot be less or equal than zero")
//	errOrderIsNil        = errors.New("order cannot be nil")
//	errSlotIsNil         = errors.New("order slot cannot be nil")
//	errResourcesIsNil    = errors.New("slot resources cannot be nil")
//	errSearchParamsIsNil = errors.New("search params cannot be nil")
//)

type Engine interface {
	ByID(ID string) (*entity.Order, error)
	Store(o *entity.Order) error
	Remove(ID string) error
}

type OrderStorage struct {
	e Engine
}

func NewOrderStorage(e Engine) *OrderStorage {
	return &OrderStorage{
		e:e,
	}
}

func (s *OrderStorage) ByID(ID string) (*entity.Order, error) {
	return s.e.ByID(ID)
}

func (s *OrderStorage) Store(o *entity.Order) error {
	return s.e.Store(o)
}

func (s *OrderStorage) Remove(ID string) error {
	return s.e.Remove(ID)
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
