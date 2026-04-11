package subnetgenerator

import "net"

// NewSubnet creates a new SubnetAddresses instance from a CIDR string.
func NewSubnet(cidr string) (*SubnetAddresses, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &SubnetAddresses{
		ipNet: ipNet,
		ip:    ip.Mask(ipNet.Mask),
	}, nil
}
