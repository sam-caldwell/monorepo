package logger

func (log *Logger[T, F]) Fatal(args ...interface{}) *Logger[T, F] {
	return log
}
