package logger

func (log *Logger) Debug(msg string) *Logger {
	if log.level >= Debug {
		log.write(Debug, &msg)
	}
	return log
}
