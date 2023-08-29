package exit

import "fmt"

// PanicOnConditionf - panic if condition is true, display formattted message
func PanicOnConditionf(condition bool, format string, msg ...any) {
	if condition {
		panic(fmt.Sprintf(format, msg...))
	}
}
