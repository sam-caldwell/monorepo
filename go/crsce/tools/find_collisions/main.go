package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"math/big"
	"time"
)

const (
	logInterval = 10 * time.Second
)

func main() {
	var objectCount uint64
	var comparisonCount uint64
	startTime := time.Now().Unix()

	empty := bytes.Repeat([]byte{0}, 1024)

	lhs := big.NewInt(0)
	rhs := big.NewInt(0)

	lhsHash := make([]byte, 1024)
	rhsHash := make([]byte, 1024)

	go func() {
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				l := hex.EncodeToString(lhs.Bytes())
				r := hex.EncodeToString(rhs.Bytes())
				log.Printf("lhs(%v) rhs(%v) "+
					"elapsed: %11.2f "+
					"objectCnt: %11d, object/sec: %11.6f "+
					"compareCnt:%11d compare/sec:%11.2f",
					l,
					r,
					elapsedTime,
					objectCount,
					float64(objectCount)/elapsedTime,
					comparisonCount,
					float64(comparisonCount)/elapsedTime)
			}
		}
	}()

	// Initialize a big.Int to cover the 1024-byte keyspace
	lhsMax := new(big.Int).Exp(big.NewInt(256), big.NewInt(1024), nil)

	for lhs = big.NewInt(17153810); lhs.Cmp(lhsMax) <= 1; lhs.Add(lhs, big.NewInt(1)) {
		objectCount++

		copy(lhsHash, empty)
		copy(lhsHash, lhs.Bytes())

		for rhs = big.NewInt(0); rhs.Cmp(lhsMax) < 0; rhs.Add(rhs, big.NewInt(1)) {
			if lhs.String() == rhs.String() {
				log.Printf("skip v:%v", rhs.String())
				continue
			}
			copy(rhsHash, empty)
			copy(rhsHash, rhs.Bytes())
			//log.Printf("value(%v, %v)", hex.EncodeToString(lhs.Bytes()), hex.EncodeToString(rhs.Bytes()))
			if bytes.Equal(lhsHash, rhsHash) {
				log.Fatalf("collision (value:%v): %v", rhs.String(), hex.EncodeToString(rhsHash))
			}
			comparisonCount++
		}
	}
}
