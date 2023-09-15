package hijack

// IteratorFunction - Function signature for use with TrackerTable.iterate()
type IteratorFunction func(key uintptr, value AppliedPatch) error
