package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	calculateSubnets "github.com/sam-caldwell/monorepo/go/net/subnetting/calculate-subnets"
	"os"
	"strconv"
)

func main() {
	exit.IfVersionRequested()
	if len(os.Args) < 3 {
		fmt.Println(calculateSubnets.ErrMissingArguments)
		os.Exit(calculateSubnets.ExitMissingArgs)
	}
	parentCIDR := os.Args[calculateSubnets.ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[calculateSubnets.ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(calculateSubnets.ExitSubnettingError)
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[calculateSubnets.ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				fmt.Println(calculateSubnets.ErrInvalidResultCount)
				os.Exit(calculateSubnets.ExitInvalidResultCount)
			}
			return int(n)
		}()
	}

	if subnets, err := calculateSubnets.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(calculateSubnets.ErrGeneral, err)
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(calculateSubnets.ExitSuccess)
}
