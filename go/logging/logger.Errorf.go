package logger

func (log *Logger) Errorf(format string, args ...any) *Logger {
	if log.level >= Error {
		log.writef(Error, &format, args...)
	}
	return log
}
