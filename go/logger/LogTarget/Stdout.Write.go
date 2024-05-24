package LogTarget

import (
	"os"
)

// Write - Write a formatted log string to stdout.
func (out Stdout[LogFormat]) Write(p *[]byte) {
	var err error
	_, err = os.Stdout.WriteString(string(*p))
	panic(err)
}
