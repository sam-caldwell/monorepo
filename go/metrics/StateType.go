package metrics

import (
	"github.com/sam-caldwell/monorepo/go/types/generic"
	hashes2 "github.com/sam-caldwell/monorepo/go/types/hashes"
)

// StateType - Define a generic StateType interface for use with StateMetric
// A StateType is essentially some type we can use to store a fixed state (e.g. a string).
type StateType interface {
	[]byte |
		string |
		hashes2.Sha1 |
		hashes2.Sha256 |
		hashes2.Sha512 |
		generic.Block1KB |
		generic.Block2KB |
		generic.Block4KB |
		generic.Block8KB |
		generic.Block16KB
}
type StateTypeMethods[T StateType] interface {
	// FromSlice - Given a byte array, store the state internally
	FromSlice([]byte)
	// ToSlice - Return a byte array from internal state
	ToSlice() []byte
	// SizeOf - Return the size of the internal structure
	SizeOf() uint
	// String - Return a string representation of the internal state
	String() string
}
