package logger

import "github.com/sam-caldwell/monorepo/go/logger/LogLevel"

// SetLevel - Define the log level
func (log *Logger[T]) SetLevel(level LogLevel.Value) *Logger[T] {
	log.level = level
	return log
}
