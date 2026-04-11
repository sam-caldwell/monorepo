package subnetgenerator

import "net"

// SubnetAddresses represents a CIDR block and provides a method to iterate through its IP addresses.
type SubnetAddresses struct {
	ipNet *net.IPNet
	ip    net.IP
}
