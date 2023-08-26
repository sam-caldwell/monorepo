package hijack

import "reflect"

// AppliedPatch - Tracking a patch to allow undo.
type AppliedPatch struct {

	// These are the original bytes we captured during hijacking...
	originalBytes []byte

	// this is a reference to the targetFunc
	target *reflect.Value

	// this is a reference to the imposterFunc
	imposter *reflect.Value
}
