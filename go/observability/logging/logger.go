package logger

import (
	"bufio"
	"os"
	"time"
)

const (
	logBufferSize = 16384 //16KB
)

type Logger struct {
	buffer        *bufio.Writer
	bufferSize    int
	err           error
	flushInterval time.Duration
	flushTimer    *time.Timer
	format        string
	level         LogLevel
	target        *os.File
}
