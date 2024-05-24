package logger

// SetAppName - Set the application name for the log messages
func (log *Logger[T]) SetAppName(appName string) *Logger[T] {
	//ToDo: sanitize app name
	log.appName = appName
	return log
}
