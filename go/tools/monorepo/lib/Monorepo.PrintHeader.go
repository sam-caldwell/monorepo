package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/terminal"
	"strings"
	"time"
)

func (m *Monorepo) PrintHeader(title string) {
	ansi.Printf("screen width: %d", terminal.GetScreenColumns()).LF().Reset()
	ansi.Cyan().
		Println(strings.Repeat("═", terminal.GetScreenColumns())).
		Printf("%s...  %v\n", title, time.Now().Format(time.RFC1123)).
		Println(strings.Repeat("═", terminal.GetScreenColumns())).
		Reset()
}
