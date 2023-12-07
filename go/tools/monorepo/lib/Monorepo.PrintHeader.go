package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/terminal/geometry"
	"time"
)

// PrintHeader - print the header text in an ascii box
func (m *Monorepo) PrintHeader(title string) {
	var window geometry.TextWindow
	width := window.ColumnSize() - 4
	geometry.DrawAsciiBox(
		ansi.Cyan(),
		[]string{
			fmt.Sprintf("%s...  %v",
				convert.Capitalizep(&title),
				time.Now().Format(time.RFC1123)),
		},
		width)
}
