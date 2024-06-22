package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/arg"
	"github.com/sam-caldwell/monorepo/go/environment"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/exit/safety"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/misc/size"
	"github.com/sam-caldwell/monorepo/go/tools/virtual-machine/machineXML"
	"github.com/sam-caldwell/monorepo/go/virtualization/kvm"
	"libvirt.org/go/libvirt"
	"log"
	"math"
	"os"
	"path/filepath"
)

func main() {
	const (
		argVmName   = "vm-name"
		argIsoFile  = "image-file"
		argRamSize  = "ram"
		argCpu      = "cpu"
		argDiskSize = "disk-size"
		argTemplate = "template"

		defaultVmName   = "myVm"
		defaultIsoFile  = "ubuntu-22.04-desktop-amd64.iso"
		defaultRam      = 1024
		defaultCpu      = 1
		defaultDiskSize = 20
		defaultTemplate = "baseline"

		minRamValue = 1
		maxRamValue = 128 * size.GB
		minCpu      = 1
		maxCpu      = 128
		minDiskSize = 20
		maxDiskSize = math.MaxUint64

		xmlBaseline      = "baseline"
		xmlUbuntuDesktop = "ubuntuDesktop"
		validVmName      = `([a-zA-Z][a-zA-Z0-9\\-\\._])*[a-zA-Z0-9]+`
	)
	var (
		err         error
		conn        *libvirt.Connect
		cpuCount    *uint
		diskPath    string
		diskSize    *uint64
		domain      *libvirt.Domain
		homeDir     string
		hostGitDir  string
		isoFileName *string
		ramSize     *uint
		vmName      *string
		vmTemplate  *string
	)

	if vmName, err = arg.String(argVmName, defaultVmName, "Vm name", validVmName); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}
	if isoFileName, err = arg.Filename(argIsoFile, defaultIsoFile, "Image filename", arg.Exists); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}
	if ramSize, err = arg.Uint(argRamSize, defaultRam, minRamValue, maxRamValue, "ram size (MB)"); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}
	if cpuCount, err = arg.Uint(argCpu, defaultCpu, minCpu, maxCpu, "number vCpu"); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}
	if diskSize, err = arg.Uint64(argDiskSize, defaultDiskSize, minDiskSize, maxDiskSize, "Disk Size (GB)"); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}
	if vmTemplate, err = arg.Choices(argTemplate, defaultTemplate, "vm template xml file", xmlBaseline, xmlUbuntuDesktop); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}
	debug := arg.Debug()

	flag.Parse()

	*diskSize = *diskSize * size.GB
	if homeDir, err = environment.RequireString("HOME"); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
	}

	// Verify/create the image directory...
	imageDir := filepath.Join(homeDir, "Documents/vm-images/")
	if !directory.Existsp(&imageDir) {
		if err = directory.Create(imageDir, 0644); err != nil {
			ansi.Red().Println(err.Error()).Fatal(exit.InvalidInput).Reset()
		}
	}
	// create the disk path...
	diskPath = filepath.Join(imageDir, fmt.Sprintf("%s.qcow2", *vmName))
	hostGitDir = filepath.Join(homeDir, "/git")

	if conn, err = libvirt.NewConnect("qemu:///system"); err != nil {
		log.Fatalf("Failed to connect to QEMU: %v", err)
	}
	safety.SafeDefer(func() { _, _ = conn.Close() })

	// Create disk image
	if _, err := os.Stat(diskPath); os.IsNotExist(err) {
		kvm.CreateDiskImage(diskPath, *diskSize)
	}

	// Define the VM XML
	var vmXML string
	switch *vmTemplate {
	case xmlUbuntuDesktop:
		vmXML = fmt.Sprintf(machineXML.XmlUbuntuDesktop, *vmName, *ramSize, *cpuCount, diskPath, *isoFileName, hostGitDir)
	default:
		vmXML = fmt.Sprintf(machineXML.XmlBaseline, *vmName, *ramSize, *cpuCount, diskPath, *isoFileName, hostGitDir)
	}

	if *debug {
		ansi.Blue().Println(vmXML).LF().Reset()
	}

	// Create the VM
	if domain, err = conn.DomainDefineXML(vmXML); err != nil {
		log.Fatalf("Failed to define domain: %v", err)
	}
	if err = kvm.CreateVm(domain); err != nil {
		log.Fatalf("Failed to create domain: %v", err)
	}

	fmt.Printf("Virtual Machine %s created and started successfully.\n", *vmName)
}
