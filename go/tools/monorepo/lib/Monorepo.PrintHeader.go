package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
	"time"
)

func (m *Monorepo) PrintHeader(title string) {
	ansi.Cyan().
		Printf("%s...  %v\n", title, time.Now().Format(time.RFC1123)).
		Println(strings.Repeat("=", ScreenWidth)).
		Reset()
}
