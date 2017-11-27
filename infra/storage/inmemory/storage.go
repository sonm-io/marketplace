package inmemory

import (
	"errors"
	"sort"
	"sync"
)

var (
	errNotFound = errors.New("element cannot be found")
)

// Storage stores and fetches elements.
type Storage struct {
	sync.RWMutex
	elements map[string]interface{}
}

// NewStorage instantiates Storage.
func NewStorage() *Storage {
	return &Storage{
		elements: make(map[string]interface{}),
	}
}

// Get fetches element by the given ID.
func (s *Storage) Get(ID string) (interface{}, error) {
	s.RLock()
	defer s.RUnlock()

	el, ok := s.elements[ID]
	if !ok {
		return nil, errNotFound
	}

	return el, nil
}

// Add adds element with the given ID to the Storage.
func (s *Storage) Add(el interface{}, ID string) error {
	s.Lock()
	defer s.Unlock()

	s.elements[ID] = el
	return nil
}

// Remove deletes element from the Storage by its ID.
// If the element is not found, an error returned.
func (s *Storage) Remove(ID string) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.elements[ID]; !ok {
		return errNotFound
	}

	delete(s.elements, ID)
	return nil
}

// Match fetches elements that match the given criteria.
func (s *Storage) Match(q ConcreteCriteria) ([]interface{}, error) {
	s.RLock()
	defer s.RUnlock()

	// the order maters, so sort the items in the elements.
	// otherwise the order is not guaranteed and may lead to errors.
	keys := make([]string, 0)
	for k := range s.elements {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var elements []interface{}
	for _, idx := range keys {
		if q.Limit > 0 && uint64(len(elements)) >= q.Limit {
			break
		}

		if q.Spec.IsSatisfiedBy(s.elements[idx]) {
			elements = append(elements, s.elements[idx])
		}
	}

	return elements, nil
}
