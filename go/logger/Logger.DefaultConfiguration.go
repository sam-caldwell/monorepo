package logger

import ratelimiter "github.com/sam-caldwell/monorepo/go/RateLimiter"

// DefaultConfiguration - Set the default values.
func (log *Logger) DefaultConfiguration() *Logger {
	log.level = DefaultLogLevel
	log.appName = DefaultLogAppName
	log.msgId = DefaultLogMsgId
	log.ratelimit = ratelimiter.NewRateLimiter(DefaultRateLimit)
	log.ConfigureStdout(nil)
	return log
}
