package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// Write - Write a formatted log string to stdout.
func (out *StdoutTarget) Write(p *[]byte) error {
	ansi.Println(string(*p))
	return nil
}
