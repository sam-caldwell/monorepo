package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Info - Write a message as a informational event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Info(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgBlue)
	if log.level.Evaluate(LogLevel.Info) {
		if _, err := log.target.Write(
			append(append(textColor,
				(&LogEvent.RFC5424Message{}).
					Create(LogLevel.Info, &log.appName, &log.msgId, &message).ToJson()...), textReset...)); err != nil {
			panic(err)
		}
	}
	return log
}
