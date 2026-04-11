package net

import "net"

// IsValidIpAddress - given an address string, return boolean true if it is a valid ipv4 or ipv6 address.
func IsValidIpAddress(address string) bool {
	return net.ParseIP(address) != nil
}
