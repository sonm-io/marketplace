package intf

// Specification represents a specification pattern.
type Specification interface {
	IsSatisfiedBy(object interface{}) bool
}

// CompositeSpecification allows chaining specifications.
type CompositeSpecification interface {
	Specification

	And(CompositeSpecification) CompositeSpecification
	Or(CompositeSpecification) CompositeSpecification
	Not() CompositeSpecification
	Relate(CompositeSpecification)
}

// BaseSpecification implements CompositeSpecification.
type BaseSpecification struct {
	CompositeSpecification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s *BaseSpecification) IsSatisfiedBy(object interface{}) bool {
	return false
}

// And creates a new AndSpecification.
func (s *BaseSpecification) And(spec CompositeSpecification) CompositeSpecification {
	a := &AndSpecification{
		s.CompositeSpecification, spec,
	}
	a.Relate(a)
	return a
}

// Or creates a new OrSpecification.
func (s *BaseSpecification) Or(spec CompositeSpecification) CompositeSpecification {
	a := &OrSpecification{
		s.CompositeSpecification, spec,
	}
	a.Relate(a)
	return a
}

// Not creates a new NotSpecification.
func (s *BaseSpecification) Not() CompositeSpecification {
	a := &NotSpecification{
		s.CompositeSpecification,
	}
	//a.Relate(a)
	return a
}

// Relate ties the given specification with its parent.
// It must only be used while creating Specifications, preferably via constructors.
func (s *BaseSpecification) Relate(spec CompositeSpecification) {
	s.CompositeSpecification = spec
}

// AndSpecification implements And condition of CompositeSpecification.
type AndSpecification struct {
	CompositeSpecification
	other CompositeSpecification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s *AndSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.CompositeSpecification.IsSatisfiedBy(object) && s.other.IsSatisfiedBy(object)
}

// OrSpecification implements Or condition of CompositeSpecification.
type OrSpecification struct {
	CompositeSpecification
	other CompositeSpecification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s *OrSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.CompositeSpecification.IsSatisfiedBy(object) || s.other.IsSatisfiedBy(object)
}

// NotSpecification implements Not condition of CompositeSpecification.
type NotSpecification struct {
	CompositeSpecification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s *NotSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.CompositeSpecification.IsSatisfiedBy(object)
}
