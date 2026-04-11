package nettools

import (
	"net"
)

// GetNetworkAddress takes a CIDR string and returns the network address.
func GetNetworkAddress(cidr string) (string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	// Get the network address from the parsed CIDR
	networkAddr := ip.Mask(ipnet.Mask).String()
	return networkAddr, nil
}
