package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Critical - Write a message as a critical event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Critical(message LogEvent.MessageValue) *Logger[T] {
	_ = log.target.Write(LogLevel.Critical, &message)
	return log
}
