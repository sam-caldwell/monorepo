package metrics

import "github.com/sam-caldwell/monorepo/go/types/generic"

// NewBigScalar initializes and returns a pointer to a new instance of BigScalar.
func NewBigScalar[ValueType generic.AnyNumber](sz uint8) *BigScalar[ValueType] {
	return &BigScalar[ValueType]{
		value:        make([]ValueType, sz),
		movingWindow: sz,
	}
}
