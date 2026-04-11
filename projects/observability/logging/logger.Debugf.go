package logger

func (log *Logger) Debugf(format string, args ...any) *Logger {
	if log.level >= Debug {
		log.writef(Debug, &format, args...)
	}
	return log
}
