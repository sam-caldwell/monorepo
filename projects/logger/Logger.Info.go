package logger

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/projects/logger/LogLevel"
)

// Info - Write a message as a informational event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Info(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgBlue)
	if log.level.Evaluate(LogLevel.Info) {
		payload := append(append(textColor,
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Info, &log.appName, &log.msgId, &log.ratelimit, &message).ToJson()...), textReset...)
		if allowed := log.ratelimit.NonBlockingCheck(len(payload)); allowed {
			if _, err := log.target.Write(payload); err != nil {
				panic(err)
			}
		}
	}
	return log
}
