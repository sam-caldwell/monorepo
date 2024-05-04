package cli

import "github.com/sam-caldwell/monorepo/go/ansi"

// SetNoop - Store the noop flag
func (client *JiraClient[T]) SetNoop(flag bool) {
	if client.debug {
		ansi.Blue().Println("noop: true").Reset()
	}
	client.noop = flag
}
