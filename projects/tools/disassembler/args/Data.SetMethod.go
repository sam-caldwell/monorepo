package args

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/exit"
)

func (data *Data) SetMethod(method *string) {
	if err := data.method.FromStringPtr(method); err != nil {
		ansi.Red().
			Println("Error: --method must specify either 'linear' or 'recursive'").
			Reset().
			Fatal(exit.InvalidInput)
	}
}
