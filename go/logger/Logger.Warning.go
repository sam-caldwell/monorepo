package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Warning - Write a message as a warning event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Warning(message LogEvent.MessageValue) *Logger {
	if log.level.Evaluate(LogLevel.Warning) {
		if _, err := log.target.Write(
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Warning, &log.appName, &log.msgId, &message).
				ToJson()); err != nil {
			panic(err)
		}
	}
	return log
}
