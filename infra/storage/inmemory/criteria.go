package inmemory


type Specification interface {
	IsSatisfiedBy(object interface{}) bool
}

type ConcreteCriteria struct {
	Limit uint64
	Spec Specification
}


