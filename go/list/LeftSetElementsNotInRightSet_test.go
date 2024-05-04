package list

import (
	"testing"
)

func TestLeftSetElementsNotInRightSet(t *testing.T) {
	t.Run("Happy path scenario", func(t *testing.T) {
		a := []any{"apple", "banana", "orange"}
		b := []any{"mango", "banana", "orange", "apple", "grape"}
		err := LeftSetElementsNotInRightSet(&a, &b)
		if err != nil {
			t.Errorf("Expected nil error, got: %v", err)
		}
	})

	t.Run("Sad path scenario", func(t *testing.T) {
		a := []any{"apple", "banana", "pear"}
		b := []any{"banana", "orange", "grape"}
		err := LeftSetElementsNotInRightSet(&a, &b)
		expectedError := "invalid element 'apple'"
		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error: %s, got: %v", expectedError, err)
		}

	})

	t.Run("Additional sad path scenario", func(t *testing.T) {
		a := []any{"apple", "banana", "orange"}
		b := []any{"apple", "banana", "orange"}
		err := LeftSetElementsNotInRightSet(&a, &b)
		if err != nil {
			t.Errorf("Expected nil error, got: %v", err)
		}
	})

}
