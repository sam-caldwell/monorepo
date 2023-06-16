package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/net/subnetting/CalculateSubnets"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(CalculateSubnets.ErrMissingArguments)
		os.Exit(CalculateSubnets.ExitMissingArgs)
	}
	parentCIDR := os.Args[CalculateSubnets.ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[CalculateSubnets.ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(CalculateSubnets.ExitSubnettingError)
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[CalculateSubnets.ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				fmt.Println(CalculateSubnets.ErrInvalidResultCount)
				os.Exit(CalculateSubnets.ExitInvalidResultCount)
			}
			return int(n)
		}()
	}

	if subnets, err := CalculateSubnets.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(CalculateSubnets.ErrGeneral, err)
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(CalculateSubnets.ExitSuccess)
}
