package monorepoCri

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/cli"
)

// List - Print list of container runtimes to stdout.
func List(filter []string) {
	cri := cli.GetCriList(filter)
	ansi.Blue().
		Line("-", 40).
		Println("Supported Container Runtimes (CRI)").
		Line("-", 40)
	for _, i := range cri {
		ansi.Printf("%v", i).LF()
	}
	ansi.Line("-", 40).LF().Reset()
}
