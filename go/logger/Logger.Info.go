package logger

func (log *Logger[T, F]) Info(args ...interface{}) *Logger[T, F] {
	return log
}
