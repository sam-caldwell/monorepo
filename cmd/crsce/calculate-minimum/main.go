package main

/*
 * cmd/crsce/calculate-minimum/main.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program calculates the minimum size (Ms).
 */

import (
	"fmt"
	"math"
)

const (
	hashSize   = 256  //sha256 hash size in bits (32*8)
	headerSize = 1216 // header size
)

func calculateMinimum(start, stop float64) (Ms float64, err error) {
	for N := start; N < stop; N++ {

		n := math.Ceil(math.Sqrt(8 * N))             // Cross Sum Size
		b := math.Ceil(math.Log2(n))                 // Cross Sum Width
		numerator := 2*b*n + hashSize*n + headerSize // Compressed Signal Size
		denominator := n * n                         // Csm Array Size.
		P := math.Abs(8*N - n*n)                     //Padding size
		diff := denominator - numerator
		fmt.Printf("numerator: %05.f  denominator: %05.f  "+
			"difference: %05.4f n: %5.2f b: %5.f N: %5.2f (P: %5.2f)\n",
			numerator, denominator, diff, n, b, N, P)
		if (P < n) && (int64(numerator) < int64(denominator)) {
			Ms = N
			return Ms, err
		}
	}
	return Ms, fmt.Errorf("ms cannot be determined")
}

func main() {
	startSz := float64(1)        //Upper search bounds
	stopSize := float64(1048576) //Lower search bounds
	Ms, err := calculateMinimum(startSz, stopSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	minimumSignalSizeInBits := Ms
	minimumSignalSizeInBytes := Ms / 8
	fmt.Printf("  minimum signal size(bytes) (Ms): %0.f\n", minimumSignalSizeInBytes)
	fmt.Printf("        minimum signal size(bits): %0.f\n", minimumSignalSizeInBits)
}
