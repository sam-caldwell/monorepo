package monorepo

import "os"

func GetCommand() *string {
	var command string
	if len(os.Args) > 1 {
		command = os.Args[1] //build,clean,test
		if len(os.Args) > 2 {
			os.Args = os.Args[1:]
		}
	}
	return &command
}
