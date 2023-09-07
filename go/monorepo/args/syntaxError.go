package args

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
)

func (arg *Arguments) syntaxError(msg string) {
	fmt.Printf("Syntax Error: %s", msg)
	os.Exit(exit.GeneralError)
}
