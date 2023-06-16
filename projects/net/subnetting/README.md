Subnetting Tools
================

## Description

A simple set of tools for working with subnets
in golang.

## Status
[![Go](https://github.com/sam-caldwell/subnetting/actions/workflows/go.yml/badge.svg)](https://github.com/sam-caldwell/subnetting/actions/workflows/go.yml)

## Commands

### calculateSubnets

> This command will calculate a list of subnet networks
> within the parent CIDR block.

**Syntax:** `calculateSubnet ${parentCIDR} ${subnetSize}`

* parentCidr = CIDR string (e.g. 10.11.0.0/16)
* subnetSize = size of the subnet (e.g. integer 0-32)
  but value must be within parent subnet.

## Package (subnetting)

### CalculateSubnets(parentCIDR string, subnetSize int) ([]string, error)

> Here's an overview of how the code works:
>
> The function takes two parameters:
>   * `parentCIDR` (a string representing the parent CIDR) and
>   * `subnetSize` (an integer representing the desired subnet size).
>
> 1. It first parses the `parentCIDR` using the net.
> 2. `ParseCIDR` function and assigns the result to `ipNet`.
> 3. If there's an error during parsing, it returns an `error`.
> 4. The function determines the number of bits set in the network mask
     > (referred to as `ones`) using the `ipNet.Mask.Size()` method.
> 5. It checks if the given `subnetSize` is valid.
> 6. If the `subnetSize` is greater than 32 or less than `ones`,
     > it returns an `error`.
> 7. The function extracts the first IP address within the parent CIDR and
     > assigns it to the `ip` variable.
> 8. It calculates the starting IP address for the subnets by combining bytes
     > from the `ip` variable.
> 9. The number of subnets within the given range is calculated using the
     > bitwise left shift operation.
> 10. A string slice called `subnets` is created with a length equal to the
      > number of subnets.
> 11. A loop iterates over each subnet and calculates the subnet IP address
      > by adding the subnet index to the starting IP.
> 12. The subnet IP address is converted back to the IP format using the
      > `fmt.Sprintf` function and assigned to the `subnet` variable.
> 13. The "subnet" string is added to the `subnets` slice.
> 14. Finally, the function returns the `subnets` slice and a nil error
      > if everything executed successfully.

