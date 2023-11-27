package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/counters"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
	"log"
	"os"
	"slices"
	"sync/atomic"
	"time"
)

const (
	defaultKeySpaceSize = 1024
	QuickTableSize      = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / ~81GB
)

func main() {
	var role string
	var genCount int
	var sortCount atomic.Int32
	var writeCount int

	hashFile := flag.String("file", "", "Pre-computed hashes")
	flag.Parse()
	log.Printf("create hash file: %v", *hashFile)

	if *hashFile == "" {
		fmt.Println("hash file (-file) required")
		os.Exit(1)
	}

	fh, err := os.Create(*hashFile)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}

	defer func() { _ = fh.Close() }()

	table := make([]hashes.Sha1, QuickTableSize)
	c, _ := counters.NewByteCounter(defaultKeySpaceSize)

	go func() {
		startTime := time.Now()
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				elapsed := time.Since(startTime)
				gProgress := 100 * float64(genCount) / float64(QuickTableSize)
				wProgress := 100 * float64(writeCount) / float64(QuickTableSize)
				sProgress := 100 * float64(sortCount.Load()) / float64(QuickTableSize)
				log.Printf("%10s (elapsed: %v) %d / %d (%8.2f %%)"+
					"write:%d (%8.2f %%) sort:%d (%8.2f %%)",
					role, elapsed, genCount, len(table), gProgress,
					writeCount, wProgress, sortCount, sProgress)
			}
		}
	}()

	role = "generating"
	for i := 0; i < QuickTableSize; i++ {
		table[i] = c.Sha1Bytes()
		_ = c.FastIncrement()
		genCount++
	}

	role = "sorting"
	sortCount.Store(0)
	slices.SortFunc(table[:], func(a, b hashes.Sha1) int {
		sortCount.Add(1)
		return bytes.Compare(a[:], b[:])
	})

	role = "writing"
	for _, hash := range table {
		if _, err := fh.WriteString(hex.EncodeToString(hash[:])); err != nil {
			panic(err)
		}
		writeCount++
	}
}
