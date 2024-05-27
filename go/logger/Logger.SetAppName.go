package logger

// SetAppName - Set the application name for the log messages
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) SetAppName(appName string) *Logger[T] {
	//ToDo: sanitize app name
	log.appName = appName
	return log
}
