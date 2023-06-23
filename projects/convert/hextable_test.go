package convert

import (
	"strings"
	"testing"
)

func TestHexTableContainsAllHexDigits(t *testing.T) {
	missingDigits := ""
	for _, digit := range "0123456789ABCDEF" {
		if !strings.ContainsRune(hexTable, digit) {
			missingDigits += string(digit)
		}
	}
	if missingDigits != "" {
		t.Errorf("hexTable is missing the following hex digits: %s", missingDigits)
	}
}
