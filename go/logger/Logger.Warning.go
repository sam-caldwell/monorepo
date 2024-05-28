package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Warning - Write a message as a warning event.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Warning(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgYellow)
	if log.level.Evaluate(LogLevel.Warning) {
		if _, err := log.target.Write(
			append(append(textColor,
				(&LogEvent.RFC5424Message{}).
					Create(LogLevel.Warning, &log.appName, &log.msgId, &message).ToJson()...), textReset...)); err != nil {
			panic(err)
		}
	}
	return log
}
