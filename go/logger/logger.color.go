package logger

import "github.com/sam-caldwell/monorepo/go/ansi"

// color - return the ANSI color codes needed for colored logs.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) color(c string) ([]byte, []byte) {
	if log.useColor {
		return []byte(c), []byte(ansi.CodeReset)
	}
	return nil, nil
}
