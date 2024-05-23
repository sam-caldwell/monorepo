package logger

func (log *Logger[T, F]) Debug(args ...interface{}) *Logger[T, F] {
	return log
}
