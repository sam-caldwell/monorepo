package logger

import (
	"github.com/sam-caldwell/monorepo/go/ratelimiter/tokensPerSecond"
)

// DefaultConfiguration - Set the default values.
func (log *Logger) DefaultConfiguration() *Logger {
	log.level = DefaultLogLevel
	log.appName = DefaultLogAppName
	log.msgId = DefaultLogMsgId
	log.ratelimit = tokensPerSecond.NewRateLimiter(DefaultRateLimit)
	log.ConfigureStdout(nil)
	return log
}
