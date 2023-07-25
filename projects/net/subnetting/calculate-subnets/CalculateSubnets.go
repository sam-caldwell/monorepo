package calculate_subnets

import (
	"fmt"
	"net"
)

// CalculateSubnets calculates subnets based on a parentCIDR and subnetSize
func CalculateSubnets(parentCIDR string, subnetSize int) (subnets []string, err error) {

	_, ipNet, err := net.ParseCIDR(parentCIDR)
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidParentCIDR, err)
	}

	ones, _ := ipNet.Mask.Size()
	if subnetSize > 32 || subnetSize < ones {
		return nil, fmt.Errorf(ErrInvalidSubnetSize, subnetSize)
	}

	// Extract the first IP address within the parent CIDR
	ip := ipNet.IP.To4()
	startIP := uint(ip[2])<<8 + uint(ip[3])

	// Calculate the number of subnets within the /16 range
	numSubnets := 1 << (subnetSize - ones)

	//subnets := make([]string, numSubnets)
	for i := 0; i < numSubnets; i++ {

		// Calculate the subnet IP by adding the subnet index to the start IP
		subnetIP := startIP + uint(i<<(32-subnetSize))

		// Convert the subnet IP back to the IP format
		subnets = append(subnets, fmt.Sprintf(MsgIpv4CIDR, ip[0], ip[1], subnetIP>>8, subnetIP&0xFF, subnetSize))

	}
	return subnets, nil
}
