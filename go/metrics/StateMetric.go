package metrics

import "sync"

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
	value []T
}

type StateType interface {
	[]byte | //arbitrary byte array
		string |
		Sha1Hash |
		Sha256Hash |
		Sha512Hash |
		Block1KB |
		Block2KB |
		Block4KB |
		Block8KB
}

type Sha1Hash [20]byte
type Sha256Hash [64]byte
type Sha512Hash [128]byte
type Block1KB [1024]byte
type Block2KB [2048]byte
type Block4KB [4096]byte
type Block8KB [8192]byte
