package types

// Sequence - Track a serial number over the lifespan of its scope
type Sequence uint64

// Get - Track
func (e *Sequence) Get() Sequence {
	*e++
	return *e
}
