package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

func (log *Logger[T, F]) Error(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Error) {
		log.target.SetLevel(logLevel.Error)
		log.target.Write(message)
		log.target.Flush()
	}
	return log
}
