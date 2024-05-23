package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
)

// Flush - Flush the log message and reset ANSI colors
func (out Stdout[LogFormat]) Flush() {
	ansi.Reset()
	_ = os.Stdout.Sync()
}
