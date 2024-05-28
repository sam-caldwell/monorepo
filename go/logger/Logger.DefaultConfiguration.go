package logger

// DefaultConfiguration - Set the default values.
func (log *Logger) DefaultConfiguration() *Logger {
	log.level = DefaultLogLevel
	log.appName = DefaultLogAppName
	log.msgId = DefaultLogMsgId
	log.rateLimit = DefaultRateLimit
	log.ConfigureStdout(nil)
	return log
}
