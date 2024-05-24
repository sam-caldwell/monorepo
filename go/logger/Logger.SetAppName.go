package logger

// SetAppName - Set the application name for the log messages
func (log *Logger[T, F]) SetAppName(appName string) {
	//ToDo: sanitize appname
	log.appName = appName
}
