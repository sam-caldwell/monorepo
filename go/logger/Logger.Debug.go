package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Debug - Write a message as a debug event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Debug(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgGreen)
	if log.level.Evaluate(LogLevel.Debug) {
		payload := append(append(textColor,
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Debug, &log.appName, &log.msgId, &log.ratelimit, &message).ToJson()...), textReset...)
		if allowed := log.ratelimit.NonBlockingCheck(len(payload)); allowed {
			if _, err := log.target.Write(payload); err != nil {
				panic(err)
			}
		}
	}
	return log
}
