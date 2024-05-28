package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Warning - Write a message as a warning event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Warning(message LogEvent.MessageValue) *Logger[T] {
	_ = log.target.Write(LogLevel.Warning, &message)
	return log
}
