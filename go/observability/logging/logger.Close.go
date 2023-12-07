package logger

func (log *Logger) Close() *Logger {
	log.err = log.target.Close()
	return log
}
