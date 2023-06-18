package lock

import "time"

const (
	createRetry     = 5
	createRetryWait = time.Millisecond * 10
)
