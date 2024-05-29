package logger

import (
	"github.com/sam-caldwell/monorepo/go/ratelimiter/tokensPerSecond"
)

// DefaultConfiguration - Set the default values.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) DefaultConfiguration() *Logger {
	log.level = DefaultLogLevel
	log.appName = DefaultLogAppName
	log.msgId = DefaultLogMsgId
	log.ratelimit = tokensPerSecond.NewRateLimiter(DefaultRateLimit)
	log.ConfigureStdout(nil)
	return log
}
