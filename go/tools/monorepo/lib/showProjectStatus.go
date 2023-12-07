package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	terminal "github.com/sam-caldwell/monorepo/go/terminal/widgets"
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
		terminal.BracketBox("enabled", terminal.Green, false).
			White().Dim().
			Printf(" (class:").
			Yellow().Bold().
			Printf("%s", *className).
			White().Dim().
			Printf(")(").
			Printf("project:").
			Yellow().Bold().
			Printf("%s", *projectName).
			White().Dim().
			Printf(" step: ").
			Yellow().Bold().
			Printf(" %s\n", *command).
			Reset()
	} else {
		terminal.BracketBox("disabled", terminal.Red, true).
			Yellow().
			White().Printf("(class:%s) (project:%s) step: %s\n",
			*className, *projectName, *command).
			Reset()
		return true
	}
	return false
}
