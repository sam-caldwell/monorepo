package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/environment"
	"testing"
)

func TestVM_Iso(t *testing.T) {

	t.Run("Initial state", func(t *testing.T) {
		var vm VM
		if vm.readOnly {
			t.Fatal("expect readOnly flag false")
		}
		if vm.errors.Size() != 0 {
			t.Fatal("expect no errors")
		}
		if i := vm.iso.Get(); i != "" {
			t.Fatal("expect vm.iso to be empty")
		}
	})

	t.Run("Test readonly flag", func(t *testing.T) {
		var vm VM
		if vm.readOnly {
			t.Fatal("expect readOnly flag false")
		}
		vm.readOnly = true
		if !vm.readOnly {
			t.Fatal("expect readOnly flag true")
		}
		if p := vm.Iso("test.iso"); p != &vm {
			t.Fatal("expect vm.Iso to return address of vm")
		}
		if vm.errors.Size() == 0 {
			t.Fatal("expect error if vm object is readonly")
		}
		if e := vm.errors.Dump(); e[0].Error() != "readonly violation (iso)" {
			t.Fatal("expected error does not match")
		}
	})

	t.Run("happy state", func(t *testing.T) {
		var vm VM
		if p := vm.Iso("{{home}}/test.iso"); p != &vm {
			t.Fatal("expect vm.Iso to return address of vm")
		}
		if vm.errors.Size() != 0 {
			t.Fatalf("expect one error got:%v iso:'%s'", vm.errors, vm.iso.Get())
		}
		envHome, err := environment.RequireString("HOME")
		if err != nil {
			t.Fatal(err)
		}
		if i := vm.iso.Get(); i != envHome+"/test.iso" {
			t.Fatal("expect vm.iso to be empty")
		}
	})

	t.Run("missing extension", func(t *testing.T) {
		var vm VM
		if p := vm.Iso("test"); p != &vm {
			t.Fatal("expect vm.Iso to return address of vm")
		}
		if vm.errors.Size() == 0 {
			t.Fatalf("expect one error got:%v iso:'%s'", vm.errors, vm.iso.Get())
		}
		if i := vm.iso.Get(); i != "" {
			t.Fatal("expect vm.iso to be empty")
		}
	})
}
