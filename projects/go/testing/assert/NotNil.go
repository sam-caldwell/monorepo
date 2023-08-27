package assert

import (
	"testing"
)

// NotNil - Print message and terminate test if entity is nil
func NotNil(t *testing.T, entity any, message string) {
	if entity == nil {
		t.Fatal(message)
	}
}
