package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/exit"
	"github.com/sam-caldwell/monorepo/go/projects/net/subnetting/calculate-subnets"
	"os"
	"strconv"
)

func main() {
	exit.IfVersionRequested()
	if len(os.Args) < 3 {
		fmt.Println(calculate_subnets.ErrMissingArguments)
		os.Exit(calculate_subnets.ExitMissingArgs)
	}
	parentCIDR := os.Args[calculate_subnets.ArgParentCIDR]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[calculate_subnets.ArgSubnetSize]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(calculate_subnets.ExitSubnettingError)
		}
		return int(n)
	}()

	//Optional result count
	resultCount := 0
	if len(os.Args) == 4 {
		resultCount = func() int {
			var err error
			var n int64
			s := os.Args[calculate_subnets.ArgResultCount]
			if n, err = strconv.ParseInt(s, 10, 32); err != nil {
				fmt.Println(calculate_subnets.ErrInvalidResultCount)
				os.Exit(calculate_subnets.ExitInvalidResultCount)
			}
			return int(n)
		}()
	}

	if subnets, err := calculate_subnets.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(calculate_subnets.ErrGeneral, err)
	} else {
		if resultCount == 0 {
			resultCount = len(subnets)
		}
		for _, network := range subnets[:resultCount] {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(calculate_subnets.ExitSuccess)
}
