package main

import (
	"encoding/hex"
	"flag"
	"github.com/sam-caldwell/monorepo/go/counters"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
	"log"
	"os"
	"time"
)

const (
	keySpaceSz = 1024
	TableSize  = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
)

func main() {
	var genCount int

	hashFile := flag.String("HashFile", "", "Pre-computed hashes")
	flag.Parse()
	if *hashFile == "" {
		log.Fatalf("hashfile required")
	}

	fh, err := os.Create(*hashFile)
	if err != nil {
		panic(err)
	}

	defer func() { _ = fh.Close() }()

	var table [TableSize]hashes.Sha1
	//table := make([]hashes.Sha1, TableSize)
	c, _ := counters.NewByteCounter(keySpaceSz)

	go func() {
		startTime := time.Now()
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				elapsed := time.Since(startTime)
				gProgress := 100 * float64(genCount) / float64(TableSize)
				gOps := float64(genCount) / float64(elapsed.Nanoseconds())
				log.Printf("generating %d / %d (%8.2f %%) %8.2f (elapsed: %vs)",
					genCount, len(table), gProgress, gOps, time.Since(startTime).Seconds())
			}
		}
	}()

	for i := 0; i < TableSize; i++ {
		table[i] = c.Sha1Bytes()
		_ = c.FastIncrement()
		genCount++
	}

	for _, hash := range table {
		_, err := fh.WriteString(hex.EncodeToString(hash[:]))
		if err != nil {
			panic(err)
		}
	}

}
