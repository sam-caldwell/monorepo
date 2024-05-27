package logger

import "github.com/sam-caldwell/monorepo/go/logger/LogLevel"

// SetLevel - Define the log level
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) SetLevel(level LogLevel.Value) *Logger[T] {
	log.target.SetLevel(level)
	return log
}
