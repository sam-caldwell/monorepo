package ordered

/*
 * projects/sets/ordered/Empty.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

// Empty - Return boolean true if set has 1 or more records
func (set *Set[T]) Empty() bool {
	return set.Count() == 0
}
