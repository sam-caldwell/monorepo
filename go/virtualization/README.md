Virtualization
==============

## Description
Create a virtualization library which reduces or eliminates the need for third-party hypervisors where the operating
system provides a native virtualization solution.

## Usage
```go 
vm:=NewVm().
	Name("myVM").
    Image("~/vm_images/{{name}}").
	Iso("~/iso/ubuntu.iso").
	Cpu(4).
	Ram(1024).
	Disk("opsysRoot",64).
	Disk("swap",1).
    NetworkInterface("eth0",NetworkBridge).
    NetworkInterface("eth1",NetworkNat).
	FailOnError().
	Create().
	Start().
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
