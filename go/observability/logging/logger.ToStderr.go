package logger

import (
	"bufio"
	"os"
)

func (log *Logger) ToStderr() *Logger {
	log.target = os.Stderr
	log.buffer = bufio.NewWriterSize(log.target, log.bufferSize)
	go log.startFlusher()
	return log
}
