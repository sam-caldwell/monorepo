package moremath

import (
	"errors"
	"fmt"
	"testing"
)

/*
 * Test for TransformLinearToRowColumn()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the linear to two-dimensional coordinate
 * transformation function.
 */

func TestTransformLinearToRowColumn(t *testing.T) {
	test := func(i, size, expectedRow, expectedColumn int, expectedError error) {
		row, column, err := TransformLinearToRowColumn(0, 1024)
		if !errors.Is(err, expectedError) {
			t.Fatalf("unexpexted error (a: %v,e: %v", err, expectedError)
		}
		if row != expectedRow {
			t.Fatalf("row is wrong (a: %d, e: %d)", row, expectedRow)
		}
		if column != expectedColumn {
			t.Fatalf("column is wrong (a: %d, e: %d)", column, expectedColumn)
		}
	}
	test(-1, 1024, 0, 0, fmt.Errorf("bounds check error"))
	test(0, -1, 0, 0, fmt.Errorf("array size error"))
	test(0, 1024, 0, 0, nil)
	test(1, 1024, 0, 1, nil)
	test(2, 1024, 0, 2, nil)
	test(3, 1024, 0, 3, nil)
	test(1023, 1024, 0, 1023, nil)
	test(1024, 1024, 1, 0, nil)
	test(2048, 1024, 2, 0, nil)
}
