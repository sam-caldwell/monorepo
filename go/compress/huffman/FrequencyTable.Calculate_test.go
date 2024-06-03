package huffman

import (
	"reflect"
	"strings"
	"testing"
)

func TestFrequencyTable_Calculate(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		input := []byte("hello world")

		expected := FrequencyTable{
			'h': 1,
			'e': 1,
			'l': 3,
			'o': 2,
			' ': 1,
			'w': 1,
			'r': 1,
			'd': 1,
		}

		var frequencyTable FrequencyTable
		frequencyTable.Calculate(input)

		if !reflect.DeepEqual(frequencyTable, expected) {
			t.Fatalf("CalculateFrequencyTable(%s) returned %v, expected %v",
				string(input), frequencyTable, expected)
		}
	})

	t.Run("High repetition test 1", func(t *testing.T) {
		input := []byte(
			strings.Repeat("a", 9) +
				strings.Repeat("b", 20) +
				strings.Repeat("c", 14) +
				strings.Repeat("d", 6))
		expected := FrequencyTable{
			'a': 9,
			'b': 20,
			'c': 14,
			'd': 6,
		}
		var frequencyTable FrequencyTable
		frequencyTable.Calculate(input)

		if !reflect.DeepEqual(frequencyTable, expected) {
			t.Fatalf("CalculateFrequencyTable(%s) returned %v, expected %v",
				string(input), frequencyTable, expected)
		}
	})

	t.Run("High repetition test 2", func(t *testing.T) {
		input := []byte("abbababbababbababbab")

		expected := FrequencyTable{
			'a': 8,
			'b': 12,
		}

		var frequencyTable FrequencyTable
		frequencyTable.Calculate(input)

		if !reflect.DeepEqual(frequencyTable, expected) {
			t.Fatalf("CalculateFrequencyTable(%s) returned %v, expected %v",
				string(input), frequencyTable, expected)
		}
	})

	t.Run("High repetition test 3", func(t *testing.T) {
		input := []byte("01010101010101010101")

		expected := FrequencyTable{
			'0': 10,
			'1': 10,
		}

		var frequencyTable FrequencyTable
		frequencyTable.Calculate(input)

		if !reflect.DeepEqual(frequencyTable, expected) {
			t.Fatalf("CalculateFrequencyTable(%s) returned %v, expected %v",
				string(input), frequencyTable, expected)
		}
	})
}
