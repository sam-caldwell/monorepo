package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/size"
	"github.com/sam-caldwell/monorepo/go/tools/virtual-machine/machineXML"
	"libvirt.org/go/libvirt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	const (
		vmName = "ubuntu-vm"
	)
	var domain *libvirt.Domain

	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("Failed to connect to QEMU: %v", err)
	}
	defer func() { _, _ = conn.Close() }()

	// Define VM parameters
	homeDir := os.Getenv("HOME")
	ram := 1024 // 1GB in MB
	vcpus := 1
	diskSize := 20 * size.GB // 20GB in bytes
	isoPath := filepath.Join(homeDir, "/Documents/iso/ubuntu-22.04.3-desktop-amd64.iso")
	diskPath := filepath.Join(homeDir, "Documents/vm-images/", fmt.Sprintf("%s.qcow2", vmName))
	hostGitDir := filepath.Join(homeDir, "/git")

	// Create disk image
	if _, err := os.Stat(diskPath); os.IsNotExist(err) {
		createDiskImage(diskPath, uint64(diskSize))
	}

	// Define the VM XML
	vmXML := fmt.Sprintf(machineXML.XmlUbuntuDesktop, vmName, ram, vcpus, diskPath, isoPath, hostGitDir)

	// Create the VM
	if domain, err = conn.DomainDefineXML(vmXML); err != nil {
		log.Fatalf("Failed to define domain: %v", err)
	}
	defer func() { _ = domain.Undefine() }()

	// Start the VM
	if err = domain.Create(); err != nil {
		log.Fatalf("Failed to create domain: %v", err)
	}
	fmt.Printf("Virtual Machine %s created and started successfully.\n", vmName)
}

func createDiskImage(path string, size uint64) {
	cmd := exec.Command("qemu-img", "create", "-f", "qcow2", path, fmt.Sprintf("%d", size))
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to create disk image: %v", err)
	}
}
