package wmiclient

import (
	"reflect"
	"testing"
	"time"
)

func TestStrToTime(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		testTime := time.Now()
		formattedTime := testTime.Format("20060102150405.000000-0700")

		var resultTime time.Time
		err := strToTime(reflect.ValueOf(&resultTime).Elem(), formattedTime)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !resultTime.Equal(testTime) {
			t.Errorf("Expected time: %v, got: %v", testTime, resultTime)
		}
	})

	t.Run("sad path: invalid type", func(t *testing.T) {
		var notATime string
		err := strToTime(reflect.ValueOf(&notATime).Elem(), "20060102150405.000000-0700")
		if err == nil {
			t.Errorf("Expected error for invalid type, got nil")
		}
	})

	t.Run("sad path: malformed string", func(t *testing.T) {
		var resultTime time.Time
		err := strToTime(reflect.ValueOf(&resultTime).Elem(), "this is not a date")
		if err == nil {
			t.Errorf("Expected error for malformed string, got nil")
		}
	})
}
