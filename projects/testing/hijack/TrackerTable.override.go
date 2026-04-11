package hijack

import (
	"reflect"
)

// override - Given a target function and imposter function, patch the imposter over the target
func (table *TrackerTable) override(target, imposter reflect.Value) (err error) {
	/*
	 *  WARNING!
	 *  This method was not exported intentionally!
	 *
	 *  You need to be sure to have checked the target and imposter to
	 *  be sure they are compatible before calling this.  If you don't
	 *  check this first, you'll have a mess worse than my second marriage.
	 *  --Sam
	 */
	/*
	 * We perform a check first to see if this method is already captured.
	 * if it is, we roll back that change before applying the next one lest
	 * we lose the original state with the second change.  Wanna guess how I
	 * learned this one was a good idea?
	 */
	if thisPatch, ok := table.patches[target.Pointer()]; ok {
		/*
		 * We call a simplified version that keeps the data in the TrackerTable
		 * but rolls back the change in memory because...it's faster than calling
		 * table.Free() or even table.rollback().  Since we are already locked,
		 * we can't call table.rollback().
		 */
		if err = rollback(target.Pointer(), thisPatch); err != nil {
			return err
		}
	}
	/*
	 * Get an address for our imposter function and call replaceFunction()
	 * to rewrite our target function (referenced by pointer) with our new
	 * hijacked code...
	 */
	imposterFuncPtr := (uintptr)(getPtr(imposter))
	var hijackedCode []byte
	if hijackedCode, err = hijackFunction(target.Pointer(), imposterFuncPtr); err != nil {
		return err
	}
	/*
	 * Add our newly applied patch to the internal patches map
	 */
	table.patches[target.Pointer()] = AppliedPatch{
		hijackedCode,
		&target,
		&imposter}
	return err
}
