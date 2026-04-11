package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/convert"
	"github.com/sam-caldwell/monorepo/projects/terminal/geometry"
	"github.com/sam-caldwell/monorepo/projects/terminal/widgets"
	"time"
)

// PrintHeader - print the header text in an ascii box
func (m *Monorepo) PrintHeader(title string) {
	var window geometry.TextWindow
	width := window.ColumnSize() - 4
	terminal.DrawAsciiBox(
		ansi.Cyan(),
		[]string{
			fmt.Sprintf("%s...  %v",
				convert.Capitalizep(&title),
				time.Now().Format(time.RFC1123)),
		},
		width)
}
