package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// Write - Write a formatted log string to stdout.
func (out Stdout) Write(p *[]byte) {
	ansi.Println(string(*p))
}
