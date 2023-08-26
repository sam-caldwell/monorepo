package hijack

import (
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"reflect"
)

// CaptureMethod - replace the method of a target with its imposter.
func (table *TrackerTable) CaptureMethod(target reflect.Type, methodName string, imposterMethod interface{}) (err error) {
	/*
	 * given a target type (target) and methodName, look up the method then replace that
	 * method with the given golang imposter function.
	 */
	method, ok := target.MethodByName(methodName)

	exit.PanicOnConditionf(!ok, "unknown method %s", methodName)

	imposter := reflect.ValueOf(imposterMethod)

	return table.override(method.Func, imposter)
}
