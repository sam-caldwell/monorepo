package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Error - Write a message as a error event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Error(message LogEvent.MessageValue) *Logger {
	if log.level.Evaluate(LogLevel.Error) {
		if _, err := log.target.Write(
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Error, &log.appName, &log.msgId, &message).
				ToJson()); err != nil {
			panic(err)
		}
	}
	return log
}
