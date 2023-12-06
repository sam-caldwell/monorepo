package monorepo

import "github.com/sam-caldwell/monorepo/go/ansi"

func showProjectStatus(enabled bool, className, projectName, command string) bool {
	ansi.White().
		Printf("  [").
		Bold()
	if enabled {
		ansi.Green().
			Printf("enabled").Reset().
			White().
			Printf("] (class:%s) (project:%s) step: %s\n", className, projectName, command).
			Reset()
	} else {
		ansi.Yellow().
			Printf("disabled").Reset().
			White().Printf("] (class:%s) (project:%s) step: %s\n", className, projectName, command).
			Reset()
		return true
	}
	return false
}
