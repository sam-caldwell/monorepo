package logger

import (
	"bufio"
	"os"
)

func (log *Logger) ToStdout() *Logger {
	log.target = os.Stdout
	log.buffer = bufio.NewWriterSize(log.target, log.bufferSize)
	go log.startFlusher()
	return log
}
