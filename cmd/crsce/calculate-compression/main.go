package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const hashSize = 256 //sha256 hash size in bits (32*8)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: missing signal size (in bytes) argument")
		os.Exit(1)
	}
	SignalSizeInBytes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: Signal Size must be an integer (in bytes)")
	}

	if SignalSizeInBytes < 8 {
		fmt.Println("Error: Signal Size must be >=8")
		os.Exit(1)
	}
	S := float64(SignalSizeInBytes) * 8 //signal size in bits

	n := math.Ceil(math.Sqrt(S)) //array size
	b := math.Ceil(math.Log2(n)) //bit width

	// This is the compression ratio of output to input signal size.
	compressionRate := (2*b*n + hashSize*n) / math.Pow(n, 2)

	fmt.Printf("results:\n  "+
		"        CompressionRate: %f  %0.00f%%\n"+
		" Signal Size (in bytes): %d\n"+
		"         Array Size (n): %f\n"+
		"    Cross Sum Width (b): %f\n",
		compressionRate, 100-(100*compressionRate),
		SignalSizeInBytes,
		n,
		b)
}
