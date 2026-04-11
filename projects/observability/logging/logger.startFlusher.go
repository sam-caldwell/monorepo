package logger

import (
	"fmt"
	"time"
)

func (log *Logger) startFlusher() {
	var defaultDuration time.Duration
	if (log.target == nil) || (log.flushInterval == defaultDuration) {
		return
	}
	log.flushTimer = time.NewTimer(log.flushInterval)
	defer log.flushTimer.Stop()
	for log.flushInterval != defaultDuration {
		select {
		case <-log.flushTimer.C:
			if err := log.buffer.Flush(); err != nil {
				panic(fmt.Sprintf("Logger Error (Flush): %v", err))
			}
		}
	}
}
