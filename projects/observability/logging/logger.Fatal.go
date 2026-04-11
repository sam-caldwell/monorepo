package logger

func (log *Logger) Fatal(msg string) *Logger {
	log.write(Fatal, &msg)
	return log
}
