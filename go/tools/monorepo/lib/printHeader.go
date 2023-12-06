package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
)

func printHeader(title string) {
	ansi.Cyan().
		Println(title).
		Println(strings.Repeat("=", ScreenWidth)).
		Reset()
}
