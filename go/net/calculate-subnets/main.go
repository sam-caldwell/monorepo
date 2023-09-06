package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	calculate_subnets2 "github.com/sam-caldwell/monorepo/go/net/subnetting/calculate-subnets"
	"os"
	"strconv"
)

func main() {
	exit.IfVersionRequested()
	if len(os.Args) < 3 {
		fmt.Println(calculate_subnets2.ErrMissingArguments)
		os.Exit(calculate_subnets2.ExitMissingArgs)
	}
	parentCIDR := os.Args[calculate_subnets2.ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[calculate_subnets2.ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(calculate_subnets2.ExitSubnettingError)
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[calculate_subnets2.ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				fmt.Println(calculate_subnets2.ErrInvalidResultCount)
				os.Exit(calculate_subnets2.ExitInvalidResultCount)
			}
			return int(n)
		}()
	}

	if subnets, err := calculate_subnets2.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(calculate_subnets2.ErrGeneral, err)
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(calculate_subnets2.ExitSuccess)
}
