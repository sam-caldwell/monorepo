package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Error - Write a message as a error event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Error(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgRed)
	if log.level.Evaluate(LogLevel.Error) {
		if _, err := log.target.Write(
			append(append(textColor,
				(&LogEvent.RFC5424Message{}).
					Create(LogLevel.Error, &log.appName, &log.msgId, &message).ToJson()...), textReset...)); err != nil {
			panic(err)
		}
	}
	return log
}
