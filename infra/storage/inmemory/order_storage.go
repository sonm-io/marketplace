package inmemory

import (
	"errors"
	"sort"
	"sync"

	// TODO: (screwyprof) break dependency
	"github.com/sonm-io/marketplace/entity"
)

var (
	errOrderNotFound = errors.New("order cannot be found")
)

// OrderStorage stores and fetches orders.
type OrderStorage struct {
	sync.RWMutex
	db map[string]*entity.Order
}

// NewStorage instantiates OrderStorage.
func NewStorage() *OrderStorage {
	return &OrderStorage{
		db: make(map[string]*entity.Order),
	}
}

// ByID fetches object and maps it into result by the given ID.
func (s *OrderStorage) ByID(ID string, result interface{}) error {
	s.RLock()
	defer s.RUnlock()

	item, ok := s.db[ID]
	if !ok {
		return errOrderNotFound
	}

	r := result.(*entity.Order)
	*r = *item


	return nil
}

// Store saves the given object.
func (s *OrderStorage) Store(o *entity.Order) error {
	s.Lock()
	defer s.Unlock()

	s.db[o.ID] = o
	return nil
}

// Remove deletes object by its ID.
func (s *OrderStorage) Remove(ID string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.db[ID]; !ok {
		return errOrderNotFound
	}

	delete(s.db, ID)
	return nil
}

// Match fetches objects that match the given criteria and maps them into result.
func (s *OrderStorage) Match(q ConcreteCriteria, result interface{}) error {
	s.RLock()
	defer s.RUnlock()


	// the order maters, so sort the items in the db.
	// otherwise the order is not guaranteed and may lead to errors.
	keys := make([]string, 0)
	for k := range s.db {
		keys = append(keys, k)
	}
	sort.Strings(keys)


	var collection []*entity.Order
	for _, idx := range keys {
		if q.Limit > 0 && uint64(len(collection)) >= q.Limit {
			break
		}

		if q.Spec.IsSatisfiedBy(s.db[idx]) {
			collection = append(collection, s.db[idx])
		}
	}

	r := result.(*[]*entity.Order)
	for idx := range collection {
		item := collection[idx]
		*r = append(*r, item)

	}

	return nil
}
