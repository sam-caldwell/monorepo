package logger

func (log *Logger[T, F]) Error(args ...interface{}) *Logger[T, F] {
	return log
}
