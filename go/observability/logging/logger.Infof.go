package logger

func (log *Logger) Infof(format string, args ...any) *Logger {
	if log.level >= Info {
		log.writef(Info, &format, args...)
	}
	return log
}
