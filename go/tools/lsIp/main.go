package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"net"
	"strings"
)

// SubnetAddresses represents a CIDR block and provides a method to iterate through its IP addresses.
type SubnetAddresses struct {
	ipnet *net.IPNet
	ip    net.IP
}

// NewSubnet creates a new SubnetAddresses instance from a CIDR string.
func NewSubnet(cidr string) (*SubnetAddresses, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &SubnetAddresses{
		ipnet: ipnet,
		ip:    ip.Mask(ipnet.Mask),
	}, nil
}

// Next returns the next IP address in the CIDR block along with a boolean indicating if there are more IPs.
func (sa *SubnetAddresses) Next() (bool, string) {
	if sa.ipnet.Contains(sa.ip) {
		nextIP := sa.ip.String()
		sa.incrementIP()
		return true, nextIP
	}
	return false, ""
}

func (sa *SubnetAddresses) incrementIP() {
	for j := len(sa.ip) - 1; j >= 0; j-- {
		sa.ip[j]++
		if sa.ip[j] > 0 {
			break
		}
	}
}

func main() {
	cidr := func() string {
		s := flag.String("cidr", "", "CIDR block")
		flag.Parse()
		if strings.TrimSpace(*s) == words.EmptyString {
			ansi.Red().Printf("CIDR block cannot be empty").LF().Fatal(1).Reset()
		}
		return strings.TrimSpace(*s)
	}()
	var err error
	var subnet *SubnetAddresses
	if subnet, err = NewSubnet(cidr); err != nil {
		ansi.Red().Printf("Error:", err).LF().Fatal(1).Reset()
	}

	ansi.Blue().
		Line("=", 40).
		Printf("Starting").LF().
		Println("IP addresses in CIDR block:").LF().
		Reset()

	for {
		hasNext, ip := subnet.Next()
		if !hasNext {
			break
		}
		ansi.Cyan().Println(ip).Reset()
	}
}
