package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"regexp"
	"testing"
)

func TestNewVm(t *testing.T) {
	vm := NewVm()
	if vm.errors.Size() != 0 {
		t.Fatalf("expect zero errors")
	}
	const namePattern = "^[a-zA-Z][a-zA-Z0-9]+$"
	regex := regexp.MustCompile(namePattern)
	if !regex.MatchString(vm.GetName()) {
		t.Fatalf("invalid name pattern")
	}
	if vm.cpuCount != 1 {
		t.Fatalf("default cpu count mismatch")
	}
	if vm.ram != 4*size.GB {
		t.Fatalf("default ram mismatch.")
	}
	if len(vm.disks) != 0 {
		t.Fatalf("expect zero disks")
	}
	if len(vm.networkInterfaces) != 0 {
		t.Fatalf("expect zero network interfaces")
	}
	if vm.image.Get() != "" {
		t.Fatalf("expect image to be empty")
	}
	if vm.iso.Get() != "" {
		t.Fatal("expect iso to be empty")
	}
	if vm.errors.Size() != 0 {
		t.Fatalf("expect zero errors")
	}
}
