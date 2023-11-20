package moremath

/*
 * TransformLinearToRowColumn()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Transform a linear address (i) to th 2D coordinate pair (r,c)
 * assuming a square matrix of (s x s) elements.
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// TransformLinearToRowColumn - Transform linear address (i) to row (r), column (c) given matrix size (s)
func TransformLinearToRowColumn(i int, s int) (row, column int, err error) {
	if i < 0 {
		return 0, 0, fmt.Errorf(errors.BoundsCheckError)
	}
	if s <= 0 {
		return 0, 0, fmt.Errorf(errors.ArraySizeError)
	}
	row = i / s
	column = i % s
	return row, column, nil
}
