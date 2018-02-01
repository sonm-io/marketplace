package sqllite

// Properties - benchmarks.
type Properties map[string]float64

// LessOrEqual returns true if Properties are less or equal to the given p2.
func (p Properties) LessOrEqual(p2 Properties) bool {
	for k, v2 := range p2 {
		// expected property is not found
		v, ok := p[k]
		if !ok {
			return false
		}

		if v > v2 {
			return false
		}
	}

	return true
}

// GreaterOrEqual returns true if Properties are greater or equal to the given p2.
func (p Properties) GreaterOrEqual(p2 Properties) bool {
	for k, v2 := range p2 {
		// expected property is not found
		v, ok := p[k]
		if !ok {
			return false
		}

		if v < v2 {
			return false
		}
	}

	return true
}
