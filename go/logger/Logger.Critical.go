package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

func (log *Logger[T, F]) Critical(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Critical) {
		log.target.SetLevel(logLevel.Critical)
		log.target.Write(message)
		log.target.Flush()
	}
	return log
}
