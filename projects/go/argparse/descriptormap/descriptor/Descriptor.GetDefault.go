package descriptor

import "fmt"

// GetDefault - Return current default
func (arg *Descriptor) GetDefault() string {
	if arg.dValue == nil {
		return "null"
	}
	return fmt.Sprintf("%v", arg.dValue)
}
