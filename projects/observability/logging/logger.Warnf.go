package logger

func (log *Logger) Warnf(format string, args ...any) *Logger {
	if log.level >= Warn {
		log.writef(Warn, &format, args...)
	}
	return log
}
