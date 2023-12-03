package logger

import "time"

func (log *Logger) Init(level LogLevel, flushInterval time.Duration) {
	log.level = level
	log.flushInterval = flushInterval
	log.ToStdout()
}
