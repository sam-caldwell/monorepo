package hijack

import "sync"

// TrackerTable - This is the table for tracking all changes we make
type TrackerTable struct {
	lock    sync.Mutex
	patches map[uintptr]AppliedPatch
}
