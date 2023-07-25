package simpleArgs

import "os"

// UseColor - Return true if the color flag is used
func UseColor() bool {
	for _, arg := range os.Args[1:] {
		if arg == "-color" || arg == "--color" {
			return true
		}
	}
	return false
}
