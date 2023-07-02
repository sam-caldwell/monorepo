package ansi

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"strings"
)

// Stats - Print the stats report
func (test *T) Stats() {
	const header = "---Test Statistics---"
	ansi.Bold().Blue().Println(header).Reset().
		Green().Printf("Pass: %d", test.pass).LF().Reset().
		Red().Printf("Fail: %d", test.fail).LF().Reset().
		Yellow().Printf("Skip: %d", test.skip).LF().Reset().
		Bold().Blue().Println(strings.Repeat(words.Hyphen, len(header))).Reset()
}
