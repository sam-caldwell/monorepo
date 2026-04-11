package logger

func (log *Logger) Tracef(format string, args ...any) *Logger {
	if log.level >= Trace {
		log.writef(Trace, &format, args...)
	}
	return log
}
