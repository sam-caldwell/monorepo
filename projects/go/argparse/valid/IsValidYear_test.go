package valid

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsValidYear(t *testing.T) {
	invalidYears := map[int]error{
		-1:    fmt.Errorf(errInvalidYear, -1),
		-2024: fmt.Errorf(errInvalidYear, -2024),
		12024: fmt.Errorf(errInvalidYear, 12024),
	}
	for year, expectedError := range invalidYears {
		expectedYear := strings.TrimSpace(fmt.Sprintf("%d", year))
		actualYear, err := IsValidYear(year)
		if err != nil {
			if err.Error() != expectedError.Error() {
				t.Fatalf("Error mismatch: '%v' '%v'", err.Error(), expectedError.Error())
			}
		} else {
			if actualYear != expectedYear {
				t.Fatalf("year (%s) does not match expectedYear (%s)", actualYear, expectedYear)
			}
		}
	}
	for year := 2023; year < 9999; year++ {
		expectedYear := strings.TrimSpace(fmt.Sprintf("%d", year))
		actualYear, err := IsValidYear(year)
		if err != nil {
			t.Fatalf("Error should be valid %d", year)
		}
		if expectedYear != actualYear {
			t.Fatalf("Expected Year %s not matching actual year %s", expectedYear, actualYear)
		}
	}
}
