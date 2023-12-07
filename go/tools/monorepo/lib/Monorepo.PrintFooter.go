package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/terminal/geometry"
	"time"
)

// PrintFooter - print the footer text in an ascii box
func (m *Monorepo) PrintFooter(command *string, err error) {
	var window geometry.TextWindow
	width := window.ColumnSize() - 2
	if err == nil {
		geometry.DrawAsciiBox(
			ansi.Green(),
			[]string{
				fmt.Sprintf("Success  %s at %v (elapsed: %vs)",
					*command,
					time.Now().Format(time.RFC1123),
					time.Since(m.StartTime).Seconds()),
			},
			width)
	} else {
		geometry.DrawAsciiBox(
			ansi.Red(),
			[]string{
				fmt.Sprintf("Failed  %s at %v (elapsed: %vs)",
					*command,
					time.Now().Format(time.RFC1123),
					time.Since(m.StartTime).Seconds()),
				fmt.Sprintf("Error: %v", err),
			},
			width)
	}
}
