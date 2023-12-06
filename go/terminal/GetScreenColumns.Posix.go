//go:build !windows
// +build !windows

package terminal

import "github.com/sam-caldwell/monorepo/go/environment"

const (
	defaultColumnWidth = 80
)

// GetScreenColumns - Return Screen width of a terminal/console (tty/pty) in columns (Posix version)
func GetScreenColumns() int {
	return environment.GetIntOrDefault("COLUMNS", defaultColumnWidth)
}
