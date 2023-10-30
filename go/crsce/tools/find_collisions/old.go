package main

import (
	"bytes"
	"crypto/sha1"
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

	lhs := big.NewInt(0)
	rhs := big.NewInt(0)

	go func() {
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				log.Printf("lhs(%s) rhs(%s) "+
					"elapsed: %11.2f "+
					"objectCnt: %11d, "+
					"object/sec: %11.2f "+
					"compareCnt:%11d "+
					"compare/sec:%11.2f",
					lhs.Text(16),
					rhs.Text(16),
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
		hash := sha1.Sum(lhs.Bytes())
		lhsHash := hash[:]

		for rhs = big.NewInt(0); rhs.Cmp(lhs) < 0; rhs.Add(rhs, big.NewInt(1)) {
			if lhs.String() == rhs.String() {
				log.Printf("skip v:%v", rhs.String())
				continue
			}
			hash := sha1.Sum(rhs.Bytes())
			rhsHash := hash[:]
			if bytes.Equal(lhsHash, rhsHash) {
				log.Fatalf("collision (value:%v): %v", rhs.String(), hex.EncodeToString(rhsHash))
			}
			comparisonCount++
		}
		objectCount += lhs.Uint64() - objectCount
	}
}
