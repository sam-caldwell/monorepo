package convert

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"math"
)

// Int32ToBytes - Write the index as a 4-byte integer
func Int32ToBytes(index int) ([]byte, error) {
	if index > math.MaxInt32 || index < math.MinInt32 {
		return nil, fmt.Errorf(errors.BoundsCheckError)
	}
	indexBytes := make([]byte, 4)
	indexBytes[0] = byte(index >> 24)
	indexBytes[1] = byte(index >> 16)
	indexBytes[2] = byte(index >> 8)
	indexBytes[3] = byte(index)
	return indexBytes, nil
}
