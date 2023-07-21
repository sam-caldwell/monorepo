package simpleArgs

import "os"

func UseColor() bool {
	for _, arg := range os.Args[1:] {
		if arg == "-color" || arg == "--color" {
			return true
		}
	}
	return false
}
