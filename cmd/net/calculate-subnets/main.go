package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	subnetting "github.com/sam-caldwell/go/v2/projects/net/subnetting/calculate-subnets"
	"os"
	"strconv"
)

func main() {
	exit.IfVersionRequested()
	if len(os.Args) < 3 {
		fmt.Println(subnetting.ErrMissingArguments)
		os.Exit(subnetting.ExitMissingArgs)
	}
	parentCIDR := os.Args[subnetting.ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[subnetting.ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(subnetting.ExitSubnettingError)
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[subnetting.ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				fmt.Println(subnetting.ErrInvalidResultCount)
				os.Exit(subnetting.ExitInvalidResultCount)
			}
			return int(n)
		}()
	}

	if subnets, err := subnetting.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(subnetting.ErrGeneral, err)
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(subnetting.ExitSuccess)
}
