package logger

import (
	"fmt"
)

func (log *Logger) write(writeLevel LogLevel, msg *string) *Logger {
	_, log.err = log.buffer.Write([]byte(fmt.Sprintf(log.format, writeLevel, *msg)))
	return log
}
