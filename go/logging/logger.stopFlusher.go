package logger

import "time"

func (log *Logger) stopFlusher() {
	var defaultDuration time.Duration
	log.flushInterval = defaultDuration
	defer log.flushTimer.Stop()
	defer log.Flush()
}
