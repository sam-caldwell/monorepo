package virtualization

import (
	"fmt"
	"testing"
)

func TestVM_GetErrors(t *testing.T) {

	var vm VM

	vm.errors.Push(fmt.Errorf("test1"))
	vm.errors.Push(fmt.Errorf("test2"))
	vm.errors.Push(fmt.Errorf("test3"))

	result := vm.GetErrors()

	if len(result) != 3 {
		t.Fatal("length of errors mismatch")
	}

	if result[0].Error() != "test1" {
		t.Fatal("test1 mismatch")
	}

	if result[1].Error() != "test2" {
		t.Fatal("test2 mismatch")
	}

	if result[2].Error() != "test3" {
		t.Fatal("test3 mismatch")
	}

}
