package moremath

/*
 * TransformLinearToRowColumn()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Transform a linear address (i) to th 2D coordinate pair (r,c)
 * assuming a square matrix of (s x s) elements.
 */
import "fmt"

// TransformLinearToRowColumn - Transform linear address (i) to row (r), column (c) given matrix size (s)
func TransformLinearToRowColumn(i, s int) (row, column int, err error) {
	if i < 0 {
		return 0, 0, fmt.Errorf("bounds check error")
	}
	if s <= 0 {
		return 0, 0, fmt.Errorf("array size error")
	}
	row = i % s
	column = i / s
	return row, column, nil
}
