package metrics

import "math/big"

type BigNumber interface {
	int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		float32 | float64 |
		big.Int | big.Float
}
