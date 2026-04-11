package logger

import (
	"fmt"
)

func (log *Logger) writef(level LogLevel, format *string, args ...any) *Logger {
	msg := fmt.Sprintf(*format, args...)
	log.write(level, &msg)
	return log
}
