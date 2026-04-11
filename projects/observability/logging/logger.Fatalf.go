package logger

func (log *Logger) Fatalf(format string, args ...any) *Logger {
	log.writef(Fatal, &format, args...)
	return log
}
