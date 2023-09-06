package ordered

/*
 * projects/sets/ordered/Set.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

import "sync"

// Set - Create map of any to its type
type Set struct {
	lock sync.RWMutex
	data []any
}
