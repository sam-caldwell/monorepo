package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Debug - Write a message as a debug event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Debug(message LogEvent.MessageValue) *Logger {
	if log.level.Evaluate(LogLevel.Debug) {
		if _, err := log.target.Write(
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Debug, &log.appName, &log.msgId, &message).
				ToJson()); err != nil {
			panic(err)
		}
	}
	return log
}
