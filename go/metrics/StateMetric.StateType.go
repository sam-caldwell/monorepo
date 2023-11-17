package metrics

// StateType - Define a generic StateType interface for use with StateMetric
// A StateType is essentially some type we can use to store a fixed state (e.g. a string).
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
type StateTypeMethods[T StateType] interface {
	FromSlice([]byte)
	ToSlice() []byte
	SizeOf() uint
}
