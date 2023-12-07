package logger

func (log *Logger) Error(msg string) *Logger {
	if log.level >= Error {
		log.write(Error, &msg)
	}
	return log
}
