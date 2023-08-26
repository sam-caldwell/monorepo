package types

import (
	"fmt"
)

// Typecheck -type-check the default value
func (arg *ArgTypes) Typecheck(value any) (err error) {
	var ok bool
	if value != nil {
		switch *arg {
		case Boolean:
			_, ok = value.(bool)
		case Flag:
			_, ok = value.(bool)
		case Float:
			_, ok = value.(float64)
		case Integer:
			_, ok = value.(int)
		case String:
			_, ok = value.(string)
		default:
			return fmt.Errorf(eMsgTypeCheckUnknown)
		}
		if !ok {
			return fmt.Errorf(eMsgTypeCheck, (*arg).String())
		}
	}
	return err
}
