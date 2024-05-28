package logger

import "github.com/sam-caldwell/monorepo/go/logger/LogLevel"

const (
	//DefaultRateLimit - Default number of log entries per minute (100/sec x 60sec = 6000)
	DefaultRateLimit = 100 * 60

	//DefaultLogLevel - This establishes the default log level
	DefaultLogLevel = LogLevel.Info

	//DefaultLogAppName - The default application name
	DefaultLogAppName = "not_set"

	//DefaultLogMsgId = Default message Id
	DefaultLogMsgId = "00"
)
