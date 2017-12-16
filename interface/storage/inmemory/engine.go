package inmemory

import "github.com/sonm-io/marketplace/infra/storage/inmemory"

// Engine represents Storage Engine.
type Engine interface {
	Get(ID string) (interface{}, error)
	Add(el interface{}, ID string) error
	Remove(ID string) error
	Match(q inmemory.ConcreteCriteria) ([]interface{}, error)
}
