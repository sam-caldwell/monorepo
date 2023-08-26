package hijack

import (
	"reflect"
)

// Capture - Patch a target function with its imposter function
func (table *TrackerTable) Capture(targetFunc, imposterFunc any) (err error) {
	/*
	 * Do you like a little Func? ...when you're drunc?
	 * Well here we go to replace the real (target) func
	 * with our patched imposter func...
	 *
	 * So to borrow from the Beastie Boys...
	 * It's time to get Func-Y!
	 */
	table.lock.Lock()
	defer table.lock.Unlock()
	/*
	 * We only capture functions which are compatible.
	 */
	if err = requireCompatibleFunctions(targetFunc, imposterFunc); err != nil {
		return err
	}
	/*
	 * The override method will take the reflected value of our target and imposter
	 * functions and perform the card trick of replacing the target with the imposter
	 * then returning any error state to the caller.
	 */
	return table.override(reflect.ValueOf(targetFunc), reflect.ValueOf(imposterFunc))
}
