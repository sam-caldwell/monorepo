package logger

// SetAppName - Set the application name for the log messages
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) SetAppName(appName string) *Logger[T] {
	if ValidAppName(&appName) {
		log.target.SetAppName(&appName)
	} else {
		panic("invalid application name")
	}
	return log
}
