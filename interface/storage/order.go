package storage

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"
	"github.com/sonm-io/marketplace/usecase/intf"
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
	ByID(ID string, result interface{}) error
	Store(o *entity.Order) error
	Remove(ID string) error
	Match(q inmemory.ConcreteCriteria, result interface{}) error
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
	var order entity.Order
	err := s.e.ByID(ID, &order)

	return &order, err
}

func (s *OrderStorage) Store(o *entity.Order) error {
	return s.e.Store(o)
}

func (s *OrderStorage) Remove(ID string) error {
	return s.e.Remove(ID)
}

func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) ([]*entity.Order, error) {

	b := inmemory.NewBuilder()
	b.WithLimit(limit)
	b.WithSpec(spec)


	var orders []*entity.Order
	err := s.e.Match(b.Build(), &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
