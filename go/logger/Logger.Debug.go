package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

func (log *Logger[T, F]) Debug(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Debug) {
		log.target.SetLevel(logLevel.Debug)
		log.target.Write(message)
		log.target.Flush()
	}
	return log
}
