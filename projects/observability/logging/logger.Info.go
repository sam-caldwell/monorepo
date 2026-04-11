package logger

func (log *Logger) Info(msg string) *Logger {
	if log.level >= Info {
		log.write(Info, &msg)
	}
	return log
}
