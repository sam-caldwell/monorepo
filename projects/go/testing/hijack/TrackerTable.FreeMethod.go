package hijack

// FreeMethod - Revert any patching done on a given method of the given target object.
//func (table *TrackerTable) FreeMethod(target reflect.Type, methodName string) bool {
//	/*
//	 * Given a target object and method name, look up the method then revert
//	 * any patches to that method (if they exist), and return a boolean value
//	 * which indicates if a revert was needed (e.g. the target had been patched).
//	 */
//	m, ok := target.MethodByName(methodName)
//
//	exit.PanicOnConditionf(!ok, "Unknown method: %s", methodName)
//
//	return unpatchValue(m.Func)
//}
