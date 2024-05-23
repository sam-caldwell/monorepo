package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
)

// File - Send log output of the given format to file
type File[LogFormat LogEvent.LogFormat] struct{}
