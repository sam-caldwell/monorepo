package args

import (
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
)

func (arg *Arguments) Parse() {
	exit.OnCondition(len(os.Args[1:]) < 3, exit.MissingArg,
		"Missing Args", checkUsage)

	arg.parameters = make(map[string]string)

	switch arg.generalOptions(os.Args[1:])[0] {
	case words.Exec:
		arg.execCommand(os.Args[1:])
	case words.Config:
		arg.config(os.Args[1:])
	case words.Build:
		arg.build(os.Args[1:])
	case words.Lint:
		arg.lint(os.Args[1:])
	case words.Test:
		arg.test(os.Args[1:])
	default:
		arg.syntaxError(
			"Expect 'exec','config','build','lint','test'," +
				"'--color','--noop','--debug'")
		os.Exit(exit.GeneralError)
	}
}
