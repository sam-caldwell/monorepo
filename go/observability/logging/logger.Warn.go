package logger

func (log *Logger) Warn(msg string) *Logger {
	if log.level >= Warn {
		log.write(Warn, &msg)
	}
	return log
}
