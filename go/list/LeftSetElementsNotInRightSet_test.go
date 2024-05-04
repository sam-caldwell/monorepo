package list

import (
	"testing"
)

func TestLeftSetElementsNotInRightSet(t *testing.T) {
	t.Run("test with []string", func(t *testing.T) {
		t.Run("Happy path scenario", func(t *testing.T) {
			a := []string{"apple", "banana", "orange"}
			b := []string{"mango", "banana", "orange", "apple", "grape"}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
		t.Run("Sad path scenario", func(t *testing.T) {
			a := []string{"apple", "banana", "pear"}
			b := []string{"banana", "orange", "grape"}
			err := LeftSetElementsNotInRightSet(&a, &b)
			expectedError := "invalid element 'apple'"
			if err == nil || err.Error() != expectedError {
				t.Errorf("Expected error: %s, got: %v", expectedError, err)
			}
		})
		t.Run("Additional sad path scenario", func(t *testing.T) {
			a := []string{"apple", "banana", "orange"}
			b := []string{"apple", "banana", "orange"}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
	})
	t.Run("test with []int", func(t *testing.T) {
		t.Run("Happy path scenario", func(t *testing.T) {
			a := []int{1, 3, 5}
			b := []int{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
		t.Run("Sad path scenario", func(t *testing.T) {
			a := []int{1, 3, 5}
			b := []int{2, 4, 6}
			err := LeftSetElementsNotInRightSet(&a, &b)
			expectedError := "invalid element '1'"
			if err == nil || err.Error() != expectedError {
				t.Errorf("Expected error: %s, got: %v", expectedError, err)
			}
		})
		t.Run("Additional sad path scenario", func(t *testing.T) {
			a := []int{1, 3, 5}
			b := []int{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
	})
	t.Run("test with []uint", func(t *testing.T) {
		t.Run("Happy path scenario", func(t *testing.T) {
			a := []uint{1, 3, 5}
			b := []uint{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
		t.Run("Sad path scenario", func(t *testing.T) {
			a := []uint{1, 3, 5}
			b := []uint{2, 4, 6}
			err := LeftSetElementsNotInRightSet(&a, &b)
			expectedError := "invalid element '1'"
			if err == nil || err.Error() != expectedError {
				t.Errorf("Expected error: %s, got: %v", expectedError, err)
			}
		})
		t.Run("Additional sad path scenario", func(t *testing.T) {
			a := []uint{1, 3, 5}
			b := []uint{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
	})
	t.Run("test with []bool", func(t *testing.T) {
		t.Run("Happy path scenario", func(t *testing.T) {
			a := []bool{true}
			b := []bool{true, false}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
		t.Run("Sad path scenario", func(t *testing.T) {
			a := []bool{true}
			b := []bool{false}
			err := LeftSetElementsNotInRightSet(&a, &b)
			expectedError := "invalid element 'true'"
			if err == nil || err.Error() != expectedError {
				t.Errorf("Expected error: %s, got: %v", expectedError, err)
			}
		})
		t.Run("Additional sad path scenario", func(t *testing.T) {
			a := []bool{false}
			b := []bool{true, false}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
	})
	t.Run("test with []float32", func(t *testing.T) {
		t.Run("Happy path scenario", func(t *testing.T) {
			a := []float32{1, 3, 5}
			b := []float32{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
		t.Run("Sad path scenario", func(t *testing.T) {
			a := []float32{1.01, 3.1415, 5.3568}
			b := []float32{2, 4, 6}
			err := LeftSetElementsNotInRightSet(&a, &b)
			expectedError := "invalid element '1.01'"
			if err == nil || err.Error() != expectedError {
				t.Errorf("Expected error: %s, got: %v", expectedError, err)
			}
		})
		t.Run("Additional sad path scenario", func(t *testing.T) {
			a := []float32{1, 3, 5}
			b := []float32{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
	})
	t.Run("test with []float64", func(t *testing.T) {
		t.Run("Happy path scenario", func(t *testing.T) {
			a := []float64{1, 3, 5}
			b := []float64{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
		t.Run("Sad path scenario", func(t *testing.T) {
			a := []float64{1.01, 3.1415, 5.3568}
			b := []float64{2, 4, 6}
			err := LeftSetElementsNotInRightSet(&a, &b)
			expectedError := "invalid element '1.01'"
			if err == nil || err.Error() != expectedError {
				t.Errorf("Expected error: %s, got: %v", expectedError, err)
			}
		})
		t.Run("Additional sad path scenario", func(t *testing.T) {
			a := []float64{1, 3, 5}
			b := []float64{1, 2, 3, 4, 5}
			err := LeftSetElementsNotInRightSet(&a, &b)
			if err != nil {
				t.Errorf("Expected nil error, got: %v", err)
			}
		})
	})

}
