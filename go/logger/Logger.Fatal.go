package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

func (log *Logger[T, F]) Fatal(message LogEvent.LogFormat) *Logger[T, F] {
	if log.level.Evaluate(logLevel.Fatal) {
		log.target.SetLevel(logLevel.Fatal)
		log.target.Write(message)
		log.target.Flush()
	}
	return log
}
