package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
)

// Http - Send log output of the given format to Http requests
type Http[LogFormat LogEvent.LogFormat] struct{}
