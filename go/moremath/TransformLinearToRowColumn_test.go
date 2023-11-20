package moremath

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
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
		t.Run(fmt.Sprintf("Test Run %d %d %d %d", i, size, expectedRow, expectedColumn), func(t *testing.T) {
			row, column, err := TransformLinearToRowColumn(i, size)
			if err == nil {
				if expectedError != nil {
					t.Fatalf("expected nil but got %v", err)
				}
			} else {
				if err.Error() != expectedError.Error() {
					t.Fatalf("expected error not found\n"+
						"actual: %v\n"+
						"expected: %v",
						err.Error(), expectedError.Error())
				}
			}
			if row != expectedRow {
				t.Fatalf("row is wrong\n"+
					"i:%d, size: %d\n"+
					"actual:   (%d,%d)\n"+
					"expected: (%d,%d)", i, size, row, column, expectedRow, expectedColumn)
			}
			if column != expectedColumn {
				t.Fatalf("column is wrong\n"+
					"actual:   (%d,%d)\n"+
					"expected: (%d,%d)", row, column, expectedRow, expectedColumn)
			}
		})
	}

	test(-1, 1024, 0, 0, fmt.Errorf(errors.BoundsCheckError))
	test(0, -1, 0, 0, fmt.Errorf(errors.ArraySizeError))
	test(0, 1024, 0, 0, nil)
	test(1, 1024, 0, 1, nil)
	test(2, 1024, 0, 2, nil)
	test(3, 1024, 0, 3, nil)
	test(1023, 1024, 0, 1023, nil)
	test(1024, 1024, 1, 0, nil)
	test(2048, 1024, 2, 0, nil)
}
