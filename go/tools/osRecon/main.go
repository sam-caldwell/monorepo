package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
)

func main() {
	exit.TerminateOnError(
		fmt.Errorf("Build error.\n" +
			"    use main.client.go        -> build client\n" +
			"    use main.server.go        -> build server\n" +
			"    use main.query-compile.go -> build query-compile\n" +
			"    use main.analyzer.go      -> build analyzer\n"))
}
