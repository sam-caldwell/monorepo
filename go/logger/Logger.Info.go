package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Info - Write a message as a informational event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Info(message LogEvent.MessageValue) *Logger[T] {
	_ = log.target.Write(LogLevel.Info, &message)
	return log
}
