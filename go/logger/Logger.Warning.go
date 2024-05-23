package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/logLevel"
)

func (log *Logger[T, F]) Warning(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Warning) {
		log.target.Write(message)
	}
	return log
}
