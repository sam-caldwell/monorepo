package args

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
)

func (data *Data) SetMethod(method *string) {
	if err := data.method.FromStringPtr(method); err != nil {
		ansi.Red().
			Println("Error: --method must specify either 'linear' or 'recursive'").
			Reset().
			Fatal(exit.InvalidInput)
	}
}
