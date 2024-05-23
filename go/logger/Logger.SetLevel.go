package logger

import "github.com/sam-caldwell/monorepo/go/logger/logLevel"

// SetLevel - Define the log level
func (log *Logger[T, F]) SetLevel(level logLevel.Value) *Logger[T, F] {
	log.level = level
	return log
}
