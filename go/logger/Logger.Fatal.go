package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"os"
)

// Fatal - Write a message as a fatal event and terminate program execution.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Fatal(message LogEvent.MessageValue) *Logger[T] {
	defer func() {
		os.Exit(1)
	}()
	_ = log.target.Write(LogLevel.Warning, &message)
	return log
}
