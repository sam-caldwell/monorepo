package metrics

import (
	"math/big"
)

// AnyNumber - A type representing any number, including the big ones
type AnyNumber interface {
	Number | big.Int | big.Float
}
