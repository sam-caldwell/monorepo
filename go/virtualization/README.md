Virtualization
==============

## Description
Create a virtualization library which reduces or eliminates the need for third-party hypervisors where the operating
system provides a native virtualization solution.

## Usage
```go 
vm:=NewVm().
	Name("myVM").
	Iso("~/iso/ubuntu.iso").
	Cpu(4).
	Ram(1024).
	Disk(0, "opsysRoot",64).
	Disk(1, "swap",1).
    NetworkInterface(0,"eth0",NetworkBridge).
    NetworkInterface(1,"eth1",NetworkNat).
	FailOnError().
	Create().
	Start().
	Image("~/vm_images/{{name}}").
	Stop().
	Destroy()
	
```

## type: VmName
> This is a datatype used by .Name() to validate the name string

```go
package main

var name VmName

if err := name.Set("myVmName"); err != nil {
    panic(err)
}

fmt.Println(name.Get())

```

## type: 