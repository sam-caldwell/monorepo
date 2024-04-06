package args

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
)

func (data *Data) OpenFile(fileType FileType, name *string, mustExist bool, flags int, permissions os.FileMode) {

	var err error
	var filePointer *os.File

	if *name == words.EmptyString {
		ansi.Red().
			Printf("Error: %s must specify an input binary file", *name).
			LF().
			Reset().
			Fatal(exit.MissingArg)
	}

	if mustExist && !file.Exists(*name) {
		ansi.Red().
			Printf("Error: '%s' file not found", *name).
			Reset().
			Fatal(exit.NotFound)
	}

	switch fileType {
	case source:
		data.source, err = os.OpenFile(*name, flags, permissions)
		filePointer = data.source
	case output:
		data.out, err = os.OpenFile(*name, flags, permissions)
		filePointer = data.out
	default:
		ansi.Red().
			Println("invalid source type").
			Reset().
			Fatal(exit.InvalidInput)
	}

	if err != nil {
		ansi.Red().
			Printf("Error: Failed to open file (%s): %v", *name, err).
			LF().
			Reset().
			Fatal(exit.FailedToOpenFile)
	}

	if data.debug {
		ansi.Cyan().
			Printf("File Information:\n").
			Printf("  Name: %s\n", filePointer.Name()).
			Printf("  Size: %d\n", file.GetFileSize(filePointer)).
			LF().
			Reset()
	}
}
