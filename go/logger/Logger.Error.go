package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Error - Write a message as a error event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Error(message LogEvent.MessageValue) *Logger[T] {
	_ = log.target.Write(LogLevel.Error, &message)
	return log
}
