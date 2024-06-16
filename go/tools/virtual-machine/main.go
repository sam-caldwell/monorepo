package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/size"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/tools/virtual-machine/machineXML"
	"libvirt.org/go/libvirt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	const (
		defaultImageFile = "ubuntu-22.04.3-desktop-amd64.iso"
		defaultRam       = 1024
		defaultCpuCount  = 1
		minimumDiskSize  = 20 * size.GB
		xmlBaseline      = "baseline"
		xmlUbuntuDesktop = "ubuntuDesktop"
	)
	var (
		err    error
		conn   *libvirt.Connect
		domain *libvirt.Domain
	)
	homeDir := os.Getenv("HOME")

	vmName := flag.String("vm-name", "", "VM name")
	imageFileName := flag.String("image-file", defaultImageFile, "Image file name")
	ramSize := flag.Uint("ram-size", defaultRam, "RAM size in MB")
	cpuCount := flag.Uint("cpu-count", defaultCpuCount, "Number virtual CPU cores")
	diskSize := flag.Uint64("disk-size", minimumDiskSize, "Disk size (in GB)")
	vmTemplate := flag.String("vmTemplate", "vmTemplate", fmt.Sprintf("vmTemplate (%s)",
		strings.Join([]string{xmlBaseline, xmlUbuntuDesktop}, words.Comma)))

	flag.Parse()
	if *vmName == "" {
		ansi.Red().Println("vm-name is required").Fatal(exit.MissingArg).Reset()
	}
	if *imageFileName == "" {
		ansi.Red().Println("image-file is required").Fatal(exit.MissingArg).Reset()
	}
	if !file.Exists(*imageFileName) {
		ansi.Red().Println("image file does not exist").Fatal(exit.MissingArg).Reset()
	}
	if *ramSize < defaultRam {
		ansi.Red().Printf("minimum ram required is %d", defaultRam).Fatal(exit.InvalidInput).Reset()
	}
	if *cpuCount < defaultCpuCount {
		ansi.Red().Printf("minimum cpu count is %d", defaultCpuCount).Fatal(exit.InvalidInput).Reset()
	}
	if *diskSize < minimumDiskSize {
		ansi.Red().Printf("minimum disk size is %d", minimumDiskSize).Fatal(exit.InvalidInput).Reset()
	}

	if conn, err = libvirt.NewConnect("qemu:///system"); err != nil {
		log.Fatalf("Failed to connect to QEMU: %v", err)
	}
	defer func() { _, _ = conn.Close() }()

	isoPath := filepath.Join(homeDir, "/Documents/iso/", *imageFileName)
	diskPath := filepath.Join(homeDir, "Documents/vm-images/", fmt.Sprintf("%s.qcow2", *vmName))
	hostGitDir := filepath.Join(homeDir, "/git")

	// Create disk image
	if _, err := os.Stat(diskPath); os.IsNotExist(err) {
		createDiskImage(diskPath, *diskSize)
	}

	// Define the VM XML
	var vmXML string
	switch *vmTemplate {
	case xmlUbuntuDesktop:
		vmXML = fmt.Sprintf(machineXML.XmlUbuntuDesktop, *vmName, *ramSize, *cpuCount, diskPath, isoPath, hostGitDir)
	default:
		vmXML = fmt.Sprintf(machineXML.XmlBaseline, *vmName, *ramSize, *cpuCount, diskPath, isoPath, hostGitDir)
	}

	// Create the VM
	if domain, err = conn.DomainDefineXML(vmXML); err != nil {
		log.Fatalf("Failed to define domain: %v", err)
	}
	defer func() { _ = domain.Undefine() }()

	// Start the VM
	if err = domain.Create(); err != nil {
		log.Fatalf("Failed to create domain: %v", err)
	}
	fmt.Printf("Virtual Machine %s created and started successfully.\n", *vmName)
}

func createDiskImage(path string, size uint64) {
	cmd := exec.Command("qemu-img", "create", "-f", "qcow2", path, fmt.Sprintf("%d", size))
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to create disk image: %v", err)
	}
}
