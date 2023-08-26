package simpleArgs

import "os"

// QuietMode - return true if the quiet flag is used.
func QuietMode() bool {
	for _, arg := range os.Args[1:] {
		if arg == "-q" || arg == "-quiet" || arg == "--quiet" {
			return true
		}
	}
	return false
}
