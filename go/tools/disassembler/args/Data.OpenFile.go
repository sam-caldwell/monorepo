package args

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

func (data *Data) OpenFile(arg string, name *string, mustExist bool, flags int, permissions os.FileMode) {
	var err error
	if *name == "" {
		ansi.Red().
			Printf("Error: %s must specify an input binary file", arg).
			LF().
			Reset().
			Fatal(exit.MissingArg)
	}

	if mustExist && !file.Exists(*name) {
		ansi.Red().
			Printf("Error: '%s' file not found", arg).
			Reset().
			Fatal(exit.NotFound)
	}
	data.source, err = os.OpenFile(*name, flags, permissions)
	if err != nil {
		ansi.Red().
			Printf("Error: Failed to open file (%s): %v", *name, err).
			LF().
			Reset().
			Fatal(exit.FailedToOpenFile)
	}
}
