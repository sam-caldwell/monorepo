package ordered

/*
 * projects/sets/ordered/List.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

// List - return a list of items
func (set *Set) List() (result []any) {
	set.lock.RLock()
	defer set.lock.RUnlock()

	if len(set.data) > 0 {
		result = make([]any, len(set.data))
		copy(result, set.data)
	}

	return result
}
