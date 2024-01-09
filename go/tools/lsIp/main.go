package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	subnetgenerator "github.com/sam-caldwell/monorepo/go/net/subnetting/generate"
	"github.com/sam-caldwell/monorepo/go/version"
	"strings"
)

func main() {
	noColor := flag.Bool("nocolor", false, "No color text")
	v := flag.Bool("version", false, "Show version")
	cidr := func() string {
		s := flag.String("cidr", "", "CIDR block")
		flag.Parse()
		return strings.TrimSpace(*s)
	}()
	if *v {
		version.Show()
	}
	if cidr == words.EmptyString {
		ansi.Red().Printf("CIDR block cannot be empty").LF().Fatal(1).Reset()
	}
	var err error
	var subnet *subnetgenerator.SubnetAddresses
	if subnet, err = subnetgenerator.NewSubnet(cidr); err != nil {
		if !*noColor {
			ansi.Red()
			defer ansi.Reset()
		}
		ansi.Printf("Error:", err).LF().Fatal(1)
	}
	if !*noColor {
		ansi.Cyan()
	}
	defer ansi.Reset()
	for {
		hasNext, ip := subnet.Next()
		if !hasNext {
			break
		}
		ansi.Println(ip)
	}
}
