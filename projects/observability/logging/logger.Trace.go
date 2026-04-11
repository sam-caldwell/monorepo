package logger

func (log *Logger) Trace(msg string) *Logger {
	if log.level >= Trace {
		log.write(Trace, &msg)
	}
	return log
}
