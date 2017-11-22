package intf

// Specification contract
type Specification interface {
	IsSatisfiedBy(object interface{}) bool
}

type CompositeSpecification interface {
	Specification

	And(Specification) CompositeSpecification
	Or(Specification) CompositeSpecification
	Not() CompositeSpecification
}


// -----------------------------------------------------------------------------

// AbstractSpecification implements CompositeSpecification interface
type AbstractSpecification struct {
	Specification
}

// And returns a specification composition with AND operator
func (c *AbstractSpecification) And(other Specification) CompositeSpecification {
	return &andSpecification{one:c, two:other}
}

// Or returns a specification composition with OR operator
func (c *AbstractSpecification) Or(other Specification) CompositeSpecification {
	return &orSpecification{one:c, two:other}
}

// Not returns a specification composition with NOT operator
func (c *AbstractSpecification) Not() CompositeSpecification {
	return &notSpecification{one:c}
}
// -----------------------------------------------------------------------------

type andSpecification struct {
	AbstractSpecification
	one Specification
	two Specification
}

func (a *andSpecification) IsSatisfiedBy(object interface{}) bool {
	return a.one.IsSatisfiedBy(object) && a.two.IsSatisfiedBy(object)
}

// -----------------------------------------------------------------------------

type orSpecification struct {
	AbstractSpecification
	one Specification
	two Specification
}

func (a *orSpecification) IsSatisfiedBy(object interface{}) bool {
	return a.one.IsSatisfiedBy(object) || a.two.IsSatisfiedBy(object)
}

// -----------------------------------------------------------------------------

type notSpecification struct {
	AbstractSpecification
	one Specification
}

func (a *notSpecification) IsSatisfiedBy(object interface{}) bool {
	return !a.one.IsSatisfiedBy(object)
}
