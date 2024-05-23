package logger

func (log *Logger[T, F]) Critical(args ...interface{}) *Logger[T, F] {
	return log
}
