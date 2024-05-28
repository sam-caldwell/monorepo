package logger

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"os"
)

// Fatal - Write a message as a fatal event and terminate program execution.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Fatal(message LogEvent.MessageValue) *Logger {
	textColor, textReset := log.color(ansi.CodeFgRed)
	if log.level.Evaluate(LogLevel.Fatal) {
		payload := append(append(textColor,
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Fatal, &log.appName, &log.msgId, &log.ratelimit, &message).ToJson()...), textReset...)
		if allowed := log.ratelimit.NonBlockingCheck(len(payload)); allowed {
			if _, err := log.target.Write(payload); err != nil {
				panic(err)
			}
		}
	}
	os.Exit(exit.GeneralError)
	return log
}
