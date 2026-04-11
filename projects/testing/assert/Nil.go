package assert

import (
	"testing"
)

// Nil - Print message and terminate test if entity is not nil
func Nil(t *testing.T, entity any, message string) {
	if entity != nil {
		t.Fatal(message)
	}
}
