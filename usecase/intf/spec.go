package intf

// Specification represents a specification pattern.
type Specification interface {
	IsSatisfiedBy(interface{}) bool
}

// CompositeSpecification allows chaining specifications.
type CompositeSpecification interface {
	Specification
	And(other Specification) CompositeSpecification
	Or(other Specification) CompositeSpecification
	Not() CompositeSpecification
}

// BaseSpecification implements CompositeSpecification.
type BaseSpecification struct {
	Specification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s BaseSpecification) IsSatisfiedBy(object interface{}) bool {
	if s.Specification == nil {
		return false
	}
	return s.Specification.IsSatisfiedBy(object)
}

// And creates a new AndSpecification.
func (s BaseSpecification) And(other Specification) CompositeSpecification {
	return BaseSpecification{AndSpecification{s, other}}
}

// Or creates a new OrSpecification.
func (s BaseSpecification) Or(other Specification) CompositeSpecification {
	return BaseSpecification{OrSpecification{s, other}}
}

// Not creates a new NotSpecification.
func (s BaseSpecification) Not() CompositeSpecification {
	return BaseSpecification{NotSpecification{s}}
}

// AndSpecification implements And condition of CompositeSpecification.
type AndSpecification struct {
	one Specification
	two Specification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s AndSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.one.IsSatisfiedBy(object) && s.two.IsSatisfiedBy(object)
}

// OrSpecification implements Or condition of CompositeSpecification.
type OrSpecification struct {
	one Specification
	two Specification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s OrSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.one.IsSatisfiedBy(object) || s.two.IsSatisfiedBy(object)
}

// NotSpecification implements Not condition of CompositeSpecification.
type NotSpecification struct {
	spec Specification
}

// IsSatisfiedBy checks if the given object satisfies the specification.
func (s NotSpecification) IsSatisfiedBy(object interface{}) bool {
	return !s.spec.IsSatisfiedBy(object)
}
