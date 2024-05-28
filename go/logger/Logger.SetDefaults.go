package logger

// SetDefaults - Set the default values.
func (log *Logger) SetDefaults() *Logger {
	log.level = DefaultLogLevel
	log.appName = DefaultLogAppName
	log.msgId = DefaultLogMsgId
	log.rateLimit = DefaultRateLimit
	return log
}
