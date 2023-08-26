package token

import (
	"testing"
)

func TestArgumentElement_GetValue(t *testing.T) {
	var input Token

	var data = []any{
		true, false,
		-1, 0, 1,
		-1.0, 0.0, 1.0,
		"test1", "",
		nil,
	}

	for _, expected := range data {
		input.value = expected
		if input.value != expected {
			t.Fatal("failed to store and retrieve expected value")
		}
		if input.GetValue() != expected {
			t.Fatal("Get() failed to retrieve expected value")
		}
	}
}
