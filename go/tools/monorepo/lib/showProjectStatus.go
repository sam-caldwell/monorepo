package monorepo

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/misc/words"
    "strings"
)

// showProjectStatus - show the project status, and if enabled, return false. if command empty, return true
func showProjectStatus(enabled bool, className, projectName, command string) bool {
    command = strings.TrimSpace(strings.TrimSuffix(command, words.NewLine))
    if command == "" {
        return true
    }
    ansi.White().
        Printf("  [").
        Bold()
    if enabled {
        ansi.Green().
            Printf("enabled").Reset().
            White().
            Printf("]  ").Dim().
            Printf("(class:").
            Yellow().Bold().
            Printf("%s", className).
            White().Dim().
            Printf(")(").
            Printf("project:").
            Yellow().Bold().
            Printf("%s", projectName).
            White().Dim().
            Printf(" step: ").
            Yellow().Bold().
            Printf(" %s\n", command).
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
