package crsce

// Crsce - Cross-Sums Compression
type Crsce struct {

	//virtual cross sum (csm) source matrix
	csm CSM
	//lateral sum and hash matrix
	lsm []CrossSum
	//vertical sum and hash matrix
	vsm []CrossSum

	//block size is the number of bits in a CSM block
	blockSize int
}
