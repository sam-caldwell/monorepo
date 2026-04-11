package logger

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/projects/logger/LogLevel"
)

// Critical - Write a message as a critical event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Critical(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgRed)
	if log.level.Evaluate(LogLevel.Critical) {
		payload := append(append(textColor,
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Critical, &log.appName, &log.msgId, &log.ratelimit, &message).ToJson()...), textReset...)
		if allowed := log.ratelimit.NonBlockingCheck(len(payload)); allowed {
			if _, err := log.target.Write(payload); err != nil {
				panic(err)
			}
		}
	}
	return log
}
