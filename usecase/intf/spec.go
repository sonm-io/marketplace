package intf

type Specification interface {
	IsSatisfiedBy(object interface{}) bool
}

type CompositeSpecification interface {
	Specification

	And(CompositeSpecification) CompositeSpecification
	Or(CompositeSpecification) CompositeSpecification
	Not() CompositeSpecification
	Relate(CompositeSpecification)
}

// -----------------------------------------------------------------------------
type BaseSpecification struct {
	CompositeSpecification
}

// Check specification
//func (s *BaseSpecification) IsSatisfiedBy(object interface{}) bool {
//	return false
//}

// Condition AND
func (s *BaseSpecification) And(spec CompositeSpecification) CompositeSpecification {
	a := &AndSpecification{
		s.CompositeSpecification, spec,
	}

	a.Relate(a)
	return a
}

// Condition OR
func (s *BaseSpecification) Or(spec CompositeSpecification) CompositeSpecification {
	a := &OrSpecification{
		s.CompositeSpecification, spec,
	}
	//a.Relate(a)
	return a
}

// Condition NOT
func (s *BaseSpecification) Not() CompositeSpecification {
	a := &NotSpecification{
		s.CompositeSpecification,
	}
	//a.Relate(a)
	return a
}

// Relate to specification
func (s *BaseSpecification) Relate(spec CompositeSpecification) {
	s.CompositeSpecification = spec
}

/////

// AndSpecification
type AndSpecification struct {
	CompositeSpecification
	other CompositeSpecification
}

// Check specification
func (s *AndSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.CompositeSpecification.IsSatisfiedBy(object) && s.other.IsSatisfiedBy(object)
}

/////

// OrSpecification
type OrSpecification struct {
	CompositeSpecification
	other CompositeSpecification
}

// Check specification
func (s *OrSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.CompositeSpecification.IsSatisfiedBy(object) || s.other.IsSatisfiedBy(object)
}

/////

// NotSpecification
type NotSpecification struct {
	CompositeSpecification
}

// Check specification
func (s *NotSpecification) IsSatisfiedBy(object interface{}) bool {
	return s.CompositeSpecification.IsSatisfiedBy(object)
}

/////
