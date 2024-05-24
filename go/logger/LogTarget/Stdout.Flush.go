package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// Flush - Flush the log message and reset ANSI colors
func (out Stdout) Flush() {
	ansi.Reset()
}
