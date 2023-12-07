package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/terminal"
	"time"
)

func (m *Monorepo) PrintHeader(title string) {
	width := terminal.GetScreenColumns() - 4
	terminal.DrawAsciiBox(
		ansi.Cyan(),
		[]string{
			fmt.Sprintf("%s...  %v",
				convert.Capitalizep(&title),
				time.Now().Format(time.RFC1123)),
		},
		width)
}
