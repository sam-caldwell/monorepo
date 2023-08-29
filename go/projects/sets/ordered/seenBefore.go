package ordered

/*
 * projects/sets/ordered/seenBefore.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

// seenBefore - returns true if an item exists
func (set *Set) seenBefore(data *any) bool {
	for _, item := range set.data {
		if item == *data {
			return true
		}
	}
	return false
}
