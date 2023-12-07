package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

func showStats(pass, fail int) {
	ansi.
		LF().
		Cyan().Printf("    Statistics\n").
		Green().Printf("      pass: %d\n", pass).
		Red().Printf("      fail: %d\n", fail).
		LF().Reset()
}
