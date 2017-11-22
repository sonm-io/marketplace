package inmemory

type CriteriaBuilder struct {
	limit   uint64
	spec Specification
}

func NewBuilder() *CriteriaBuilder {
	return &CriteriaBuilder{}
}

func (b *CriteriaBuilder) WithLimit(limit uint64) {
	b.limit = limit
}

func (b *CriteriaBuilder) WithSpec(spec Specification) {
	if spec == nil {
		return
	}
	b.spec = spec
}

func (b *CriteriaBuilder) Build() ConcreteCriteria {
	c := ConcreteCriteria{}
	c.Spec = b.spec
	c.Limit = b.limit

	return c
}
