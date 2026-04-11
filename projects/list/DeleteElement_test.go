package list

import (
	"reflect"
	"testing"
)

func TestDeleteElement(t *testing.T) {
	// Happy path scenario
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	expectedResult := []int{1, 2, 4, 5}
	result := DeleteElement(slice, index)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Expected result: %v, got: %v", expectedResult, result)
	}

	// Sad path scenario (index out of range)
	index = 10
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected a panic due to index out of range, but no panic occurred")
		}
	}()
	_ = DeleteElement(slice, index)
}
