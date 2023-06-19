package simple

import (
	"github.com/sam-caldwell/go/v2/projects/misc"
	"testing"
)

func TestSet_TypeCheck(t *testing.T) {
	var set Set

	set.Init()

	set.data[1] = misc.NullObjectStruct{}

	if !set.TypeCheck(2) {
		t.Fail()
	}
	if set.TypeCheck(true) {
		t.Fail()
	}
	if set.TypeCheck(1.0) {
		t.Fail()
	}
	if set.TypeCheck(false) {
		t.Fail()
	}
	if set.TypeCheck("1") {
		t.Fail()
	}
}
