package inmemory

// CriteriaBuilder helps building arbitrary Criteria objects.
type CriteriaBuilder struct {
	limit uint64
	spec  Specification
}

// NewBuilder creates a new instance of CriteriaBuilder.
func NewBuilder() *CriteriaBuilder {
	return &CriteriaBuilder{}
}

// WithLimit adds a limit to the Criteria.
func (b *CriteriaBuilder) WithLimit(limit uint64) {
	b.limit = limit
}

// WithSpec adds a Specification to the Criteria.
func (b *CriteriaBuilder) WithSpec(spec Specification) {
	if spec == nil {
		return
	}
	b.spec = spec
}

// Build builds the Criteria object.
func (b *CriteriaBuilder) Build() ConcreteCriteria {
	c := ConcreteCriteria{}
	c.Spec = b.spec
	c.Limit = b.limit

	return c
}
