package metrics

import (
	"sync"
)

/*
 * Theory of Use
 *	The StateMetric is intended to store a single string of bytes or similar
 *  state.  This is often useful when showing a sample of state, and you are
 *  interested in showing the state or measuring its change.
 *
 *  Examples:
 * 		* hashes or strings (where you measure the change)
 *      * simple byte-string state just to summarize the current state of things
 */

// StateMetric - A single metric capable of storing a number of state values (in buckets) for analysis
type StateMetric[T StateType] struct {
	lock  sync.Mutex
	value T
}
