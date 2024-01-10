package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc"
	nettools "github.com/sam-caldwell/monorepo/go/net/getNetworkAddress"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

func main() {
	cidr := flag.String("cidr", "", "cidr block")
	trim := flag.Bool("trim", false, "optional flag to trim right hand .0 octets")
	ver := flag.Bool("version", false, "show version")
	flag.Parse()
	if *ver {
		version.Show()
	}
	if *cidr == "" {
		ansi.Red().Println("invalid cidr block").Fatal(1).Reset()
	}
	networkAddr, err := nettools.GetNetworkAddress(*cidr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if *trim {
		fmt.Println(misc.TrimAllSuffix(networkAddr, ".0"))
	} else {
		fmt.Println(networkAddr)
	}
}
