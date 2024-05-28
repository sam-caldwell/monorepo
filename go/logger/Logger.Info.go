package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Info - Write a message as a informational event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Info(message LogEvent.MessageValue) *Logger {
	if log.level.Evaluate(LogLevel.Info) {
		if _, err := log.target.Write(
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Info, &log.appName, &log.msgId, &message).
				ToJson()); err != nil {
			panic(err)
		}
	}
	return log
}
