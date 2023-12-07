package logger

func (log *Logger) Flush() *Logger {
	if log.target != nil {
		_ = log.buffer.Flush()
	}
	return log
}
