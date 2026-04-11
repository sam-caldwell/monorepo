package hijack

// FreeAll removes all applied monkey patches
//func (table *TrackerTable) FreeAll() (err error) {
//	table.lock.Lock()
//	defer table.lock.Unlock()
//
//	return table.iterate(func(key uintptr, value AppliedPatch) error {
//		unpatch(key, value)
//		return table.delete(key)
//	})
//}
