package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

func (log *Logger[T, F]) Info(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Info) {
		log.target.SetLevel(logLevel.Info)
		log.target.Write(message)
		log.target.Flush()
	}
	return log
}
