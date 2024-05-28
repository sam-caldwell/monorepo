package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Critical - Write a message as a critical event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Critical(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgRed)
	if log.level.Evaluate(LogLevel.Critical) {
		if _, err := log.target.Write(
			append(append(textColor,
				(&LogEvent.RFC5424Message{}).
					Create(LogLevel.Critical, &log.appName, &log.msgId, &message).ToJson()...), textReset...)); err != nil {
			panic(err)
		}
	}
	return log
}
