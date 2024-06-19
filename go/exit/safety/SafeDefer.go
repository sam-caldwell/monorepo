package safety

import "sync"

var (
	defers []func()
	mu     sync.Mutex
)

// Defer - registers a deferred function with our FIFO queue of functions to call
//
//	(c) 2023 Sam Caldwell.  MIT License
func Defer(df func()) {
	mu.Lock()
	defer mu.Unlock()
	defers = append(defers, df)
}

// ExecuteAll runs all registered deferred functions in FIFO order.
//
//	(c) 2023 Sam Caldwell.  MIT License
func ExecuteAll() {
	mu.Lock()
	defer mu.Unlock()
	for _, df := range defers {
		df()
	}
	defers = nil // Clear the slice after executing.
}

// Automatically call ExecuteAll at the end of the program
//
//	(c) 2023 Sam Caldwell.  MIT License
func init() {
	defer ExecuteAll()
}
