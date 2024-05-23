package LogTarget

import (
	"os"
)

// WriteString - Write a formatted log string to stdout.
func (out *Stdout[LogFormat]) WriteString(p LogFormat) (err error) {
	var buf []byte
	buf, err = p.ToJson()
	if err != nil {
		return err
	}
	_, err = os.Stdout.WriteString(string(buf))
	return err
}
