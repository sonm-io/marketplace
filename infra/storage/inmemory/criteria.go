package inmemory

// Specification represents a Specification pattern.
type Specification interface {
	IsSatisfiedBy(object interface{}) bool
}

// ConcreteCriteria is used a parameter object to encapsulate Criteria to match against.
type ConcreteCriteria struct {
	Limit uint64
	Spec  Specification
}
