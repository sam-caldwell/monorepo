package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
)

func printFooter() {
	ansi.Green().LF().Println(strings.Repeat("=", ScreenWidth)).Reset()
}
