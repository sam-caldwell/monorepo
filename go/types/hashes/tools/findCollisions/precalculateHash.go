package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"github.com/sam-caldwell/monorepo/go/counters"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
	"log"
	"os"
	"slices"
	"time"
)

const (
	keySpaceSz = 1024
	TableSize  = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
)

func main() {
	var genCount int
	var writeCount int

	hashFile := flag.String("file", "hashes.txt", "Pre-computed hashes")
	flag.Parse()
	log.Printf("create hash file: %v", *hashFile)

	if *hashFile == "" {
		log.Fatalf("hashfile required")
	}

	fh, err := os.Create(*hashFile)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}

	defer func() { _ = fh.Close() }()

	table := make([]hashes.Sha1, TableSize)
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
				wProgress := 100 * float64(writeCount) / float64(TableSize)
				gOps := float64(genCount) / float64(elapsed.Nanoseconds())
				log.Printf("generating %d / %d (%8.2f %%) %8.2f (elapsed: %vs) writeCount:%d (%8.2f %%)",
					genCount, len(table), gProgress, gOps, time.Since(startTime).Seconds(), writeCount, wProgress)
			}
		}
	}()

	for i := 0; i < TableSize; i++ {
		table[i] = c.Sha1Bytes()
		_ = c.FastIncrement()
		genCount++
	}

	slices.SortFunc(table[:], func(a, b hashes.Sha1) int {
		return bytes.Compare(a[:], b[:])
	})

	for _, hash := range table {
		if _, err := fh.WriteString(hex.EncodeToString(hash[:])); err != nil {
			panic(err)
		}
	}
}
