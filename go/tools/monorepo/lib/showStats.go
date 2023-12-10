package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

func showStats(pass, fail int) {
	ansi.
		LF().
		Cyan().Printf("    Statistics\n").
		Green().Printf("      pass: %d\n", pass)
	if fail > 0 {
		ansi.Red().Printf("      fail: %d\n", fail)
	}
	ansi.LF().Reset()
}
