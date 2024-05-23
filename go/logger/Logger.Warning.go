package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

func (log *Logger[T, F]) Warning(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Warning) {
		log.target.SetLevel(logLevel.Warning)
		log.target.Write(message)
		log.target.Flush()
	}
	return log
}
