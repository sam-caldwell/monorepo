package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// showProjectStatus - show the project status, and if enabled, return false. if command empty, return true
func showProjectStatus(enabled bool, className, projectName, command *string) bool {
	*command = strings.TrimSpace(strings.TrimSuffix(*command, words.NewLine))
	if *command == "" {
		return true
	}
	//ansi.White().
	//	Printf("  [").
	//	Bold()
	if enabled {
		ansi.
			White().Dim().
			Printf("     (class:").
			Yellow().Bold().
			Printf("%s", *className).
			White().Dim().
			Printf(")(").
			Printf("project:").
			Yellow().Bold().
			Printf("%s", *projectName).
			White().Dim().
			Printf(")").
			Yellow().Bold().LF().
			Reset()
	} else {
		ansi.
			Yellow().
			White().Printf("(class:%s) (project:%s)", *className, *projectName).
			Reset().LF()
		return true
	}
	return false
}
