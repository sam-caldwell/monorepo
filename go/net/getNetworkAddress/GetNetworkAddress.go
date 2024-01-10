package nettools

import (
	"github.com/sam-caldwell/monorepo/go/misc"
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

func GetNetworkAddressTrimmed(cidr string) (string, error) {
	addr, err := GetNetworkAddress(cidr)
	if err != nil {
		return "", err
	}
	return misc.TrimAllSuffix(addr, ".0"), nil
}
