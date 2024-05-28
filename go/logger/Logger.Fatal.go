package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Fatal - Write a message as a fatal event and terminate program execution.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Fatal(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgRed)
	if log.level.Evaluate(LogLevel.Fatal) {
		if _, err := log.target.Write(
			append(append(textColor,
				(&LogEvent.RFC5424Message{}).
					Create(LogLevel.Fatal, &log.appName, &log.msgId, &message).ToJson()...), textReset...)); err != nil {
			panic(err)
		}
	}
	return log
}
