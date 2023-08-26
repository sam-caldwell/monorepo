package exit

// OnError - A standard way to safely terminate on error
func OnError(err error, code int, usage string) {
	if err != nil {
		OnCondition(true, code, err.Error(), usage)
	}
}
