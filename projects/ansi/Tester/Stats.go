package ansi

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"strings"
)

// Stats - Print the stats report
func (test *T) Stats() {
	const header = "---Test Statistics---"
	ansi.Bold().Blue().Println(header).
		Printf("Pass: %d", test.pass).LF().
		Printf("Fail: %d", test.fail).LF().
		Printf("Skip: %d", test.skip).LF().
		Bold().Println(strings.Repeat("-", len(header))).LF().Reset()
}
