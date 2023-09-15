package ordered

/*
 * projects/sets/ordered/Delete.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

func (set *Set) Delete(pos int) error {
	if (pos >= 0) && (pos < len(set.data)) {
		set.lock.Lock()
		defer set.lock.Unlock()
		set.data = append(set.data[:pos], set.data[pos+1:]...)
		return nil
	}
	return fmt.Errorf(errors.IndexOutOfRange)
}
