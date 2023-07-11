package exit

// PanicOnCondition - panic if condition is true, display message
func PanicOnCondition(condition bool, msg any) {
	if condition {
		panic(msg)
	}
}
