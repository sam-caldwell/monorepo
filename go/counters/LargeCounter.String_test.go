package counters

/*
 * LargeCounter.String() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the String() method
 */

import "testing"

func TestLargeCounterString(t *testing.T) {
	func() {
		// Create a LargeCounter with some sample data
		counter := LargeCounter{0x007b, 0x01c8, 0x0315}

		// Calculate the expected hexadecimal string
		expected := "" +
			"00" + "00" + "00" + "00" + "00" + "00" + "03" + "15" +
			"00" + "00" + "00" + "00" + "00" + "00" + "01" + "c8" +
			"00" + "00" + "00" + "00" + "00" + "00" + "00" + "7b"

		// Get the actual hexadecimal string
		actual := counter.String()

		// Check if the actual result matches the expected result
		if actual != expected {
			t.Errorf("String() comparsion(1) failed\n"+
				"actual:  %s,\n"+
				"expected %s",
				actual, expected)
		}
	}()
	func() {
		counter := LargeCounter{0, 0, 0}
		expected := "" +
			"00" + "00" + "00" + "00" + "00" + "00" + "00" + "00" +
			"00" + "00" + "00" + "00" + "00" + "00" + "00" + "00" +
			"00" + "00" + "00" + "00" + "00" + "00" + "00" + "00"

		actual := counter.String()
		// Check if the actual result matches the expected result
		if actual != expected {
			t.Errorf("String() comparsion(2) failed\n"+
				"actual:  %s,\n"+
				"expected %s",
				actual, expected)
		}
	}()
}
