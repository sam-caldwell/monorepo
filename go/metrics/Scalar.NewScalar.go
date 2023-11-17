package metrics

import "github.com/sam-caldwell/monorepo/go/types/generic"

// NewScalar initializes and returns a pointer to a new instance of Scalar.
func NewScalar[ValueType generic.AnyNumber, CounterType generic.AnyNumber](sz uint8) *Scalar[ValueType, CounterType] {
	return &Scalar[ValueType, CounterType]{
		value: make([]ValueType, sz),
		//count:        0, ...this is implied
		movingWindow: sz,
	}
}
