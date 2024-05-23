package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
)

// Syslog - Send log output of the given format to Syslog
type Syslog[LogFormat LogEvent.LogFormat] struct{}
