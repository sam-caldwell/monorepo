package cli

import "github.com/sam-caldwell/monorepo/go/ansi"

// SetDebug - Store the debug flag
func (client *JiraClient[T]) SetDebug(flag bool) {
    if flag {
        ansi.Blue().Println("debug: true").Reset()
    }
    client.debug = flag
}
