package hijack

import "reflect"

// unpatchValue - remove patch from target function, return bool indicating if there had been a patch applied.
func (table *TrackerTable) unpatchValue(target reflect.Value) (ok bool, err error) {
	table.lock.Lock()
	defer table.lock.Unlock()
	if thisPatch, ok := table.patches[target.Pointer()]; ok {
		if err = rollback(target.Pointer(), thisPatch); err != nil {
			return false, err
		}
		err = table.delete(target.Pointer())
	} else {
		return false, err
	}

	return true, err
}
