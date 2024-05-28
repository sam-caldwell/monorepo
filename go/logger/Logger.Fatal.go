package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"os"
)

// Fatal - Write a message as a fatal event and terminate program execution.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Fatal(message LogEvent.MessageValue) *Logger {
	if log.level.Evaluate(LogLevel.Fatal) {
		if _, err := log.target.Write(
			(&LogEvent.RFC5424Message{}).
				Create(LogLevel.Fatal, &log.appName, &log.msgId, &message).
				ToJson()); err != nil {
			panic(err)
		}
	}
	os.Exit(1)
	return log
}
