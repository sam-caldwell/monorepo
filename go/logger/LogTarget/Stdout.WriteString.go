package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"os"
)

// WriteString - Write a formatted log string to stdout.
func (out Stdout[LogFormat]) WriteString(p LogEvent.LogFormat) {
	var buf []byte
	var err error
	buf, err = p.ToJson()
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.WriteString(string(buf))
	panic(err)
}
