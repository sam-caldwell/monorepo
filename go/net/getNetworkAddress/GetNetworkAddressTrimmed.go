package nettools

import "github.com/sam-caldwell/monorepo/go/misc"

// GetNetworkAddressTrimmed - Given a CIDR address, trim the right-hand zero (0) octets
func GetNetworkAddressTrimmed(cidr string) (string, error) {
	addr, err := GetNetworkAddress(cidr)
	if err != nil {
		return "", err
	}
	return misc.TrimAllSuffix(addr, ".0"), nil
}
