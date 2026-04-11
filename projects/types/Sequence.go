package types

// Sequence - Track a serial number over the lifespan of its scope
type Sequence uint64

// Get - Track
func (e *Sequence) Get() Sequence {
	*e++
	return *e
}

func (e *Sequence) Current() uint64 {
	return uint64(*e)
}
