package counters

/*
 * LargeCounter.Bytes() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the Bytes() method
 */

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLargeCounter_Bytes(t *testing.T) {
	func() {
		// Create a LargeCounter with some sample data
		counter := LargeCounter{123, 456, 789}

		// Calculate the expected byte representation in big-endian order
		expected := []byte{
			0, 0, 0, 0, 0, 0, 3, 21, // 789 in big-endian
			0, 0, 0, 0, 0, 0, 1, 200, // 456 in big-endian
			0, 0, 0, 0, 0, 0, 0, 123, // 123 in big-endian
		}

		// Get the actual byte representation
		actual := counter.Bytes()

		// Check if the actual result matches the expected result

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Bytes() returned\n"+
				"actual   %v,\n"+
				"expected %v",
				actual, expected)
		}
	}()

	func() {
		counter := LargeCounter{0, 0, 0}
		actual := counter.Bytes()
		expected := []byte{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("expected byte value is wrong(1):(\n"+
				"actual:  %v,\n"+
				"expected:%v)",
				actual, expected)
		}
		counter.Increment()
		if counter[0] != 1 {
			t.Fatal("incremented wrong position")
		}
		actual = counter.Bytes()
		expected = []byte{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 1,
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("expected byte value is wrong(2):(\n"+
				"actual:%v\n,"+
				"expected:%v)",
				actual, expected)
		}
		expected = []byte{
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 1,
		}
		actual = counter.Bytes()
		if !bytes.Equal(actual, expected) {
			t.Fatalf("expected byte order does not match\n"+
				"actual:   %v\n"+
				"expected: %v",
				actual, expected)
		}
	}()
}
