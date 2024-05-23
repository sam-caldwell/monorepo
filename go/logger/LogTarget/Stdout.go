package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
)

// Stdout - Send log output of the given format to stdout
type Stdout[LogFormat LogEvent.LogFormat] struct{}
