package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Debug - Write a message as a debug event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Debug(message LogEvent.MessageValue) *Logger[T] {
	_ = log.target.Write(LogLevel.Debug, &message)
	return log
}
