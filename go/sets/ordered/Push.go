package ordered

/*
 * projects/sets/ordered/Push.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

// Push - Push the item to the top of the list (n+1)
func (set *Set) Push(item any) error {
	return set.Add(item)
}
