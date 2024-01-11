package nettools

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc"
	"strconv"
	"strings"
)

// GetNetworkAddressTrimmed - Given a CIDR address, trim the right-hand zero (0) octets
func GetNetworkAddressTrimmed(cidr string) (string, error) {
	addr, err := GetNetworkAddress(cidr)
	if err != nil {
		return "", err
	}
	parts := strings.Split(cidr, "/")
	sz, err := func() (int, error) {

		if len(parts) < 2 {
			return 0, fmt.Errorf("invalid cidr address")
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		return n / 8, nil
	}()

	result := func() string {
		return misc.TrimSuffixCount(addr, ".0", sz)
	}()

	return result, nil
}
