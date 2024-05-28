package logger

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// SetLevel - Define the log level
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) SetLevel(level LogLevel.Value) *Logger {
	switch level {
	case LogLevel.Debug, LogLevel.Info, LogLevel.Warning, LogLevel.Error, LogLevel.Critical, LogLevel.Fatal:
		log.level = level
	default:
		panic(errors.InvalidLogLevel)
	}
	return log
}
