package metrics

import "github.com/sam-caldwell/monorepo/go/types/generic"

// NewScalar initializes and returns a pointer to a new instance of Scalar.
func NewScalar[ValueType generic.AnyNumber](sz uint8) *Scalar[ValueType] {
	return &Scalar[ValueType]{
		value: make([]ValueType, sz),
		//count:        0, ...this is implied
		movingWindow: sz,
	}
}
