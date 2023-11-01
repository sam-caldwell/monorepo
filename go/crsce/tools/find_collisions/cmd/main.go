package main

/*
 * FindCollisions Command
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a command to scan the keyspace of a given size
 * across a given number of concurrent workers.
 */

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	logInterval         = 5 * time.Second
	defaultKeySpaceSize = 1024
	candidateQueueSize  = 65536
)

type Candidate struct {
	raw  string
	hash string
}

func main() {
	keySpaceSize := flag.Uint(
		"keySpaceSize",
		defaultKeySpaceSize,
		"Number of bytes in the key space to scan")

	NumberOfWorkers := flag.Uint(
		"numberWorkers",
		uint(runtime.NumCPU()),
		"Number of workers to launch")

	flag.Parse()

	if *NumberOfWorkers > 255 {
		log.Println("number of workers exceeds max worker count")
	}

	log.Printf("Starting with %d generator workers (keySpaceSz:%d)\n", *NumberOfWorkers, *keySpaceSize)

	startTime := time.Now().Unix()
	count := uint64(0)
	queue := make(chan Candidate, candidateQueueSize)
	completedPasses := 0

	rhsWorkerCount := 0

	go func() {
		log.Printf("Log interval: %d", logInterval)
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				chgCompletedPasses := float64(completedPasses) / float64(elapsedTime)

				log.Printf("elapsed: %f objectCnt: %d, object/sec: %6.2f "+
					"queueSz:%d rhsWorkerCount:%d completedPasses: %6.2f",
					elapsedTime, count, float64(count)/elapsedTime, len(queue), rhsWorkerCount, chgCompletedPasses)
			}
		}
	}()
	numLhsWorkers := 1
	if *NumberOfWorkers > 1 {
		numLhsWorkers = int(*NumberOfWorkers) / 2
	}
	if *NumberOfWorkers > 8 {
		numLhsWorkers = int(*NumberOfWorkers) / 4
	}
	for i := 0; i < numLhsWorkers; i++ {
		go func(offset int) {
			c, _ := counters.NewByteCounter(int(*keySpaceSize))
			_ = c.Set(0, byte(offset))
			for {
				queue <- Candidate{
					raw:  c.String(),
					hash: c.Sha1(),
				}
				if err := c.Add(numLhsWorkers); err != nil {
					return
				}
			}
		}(i)
		log.Printf("worker %d started", i)
	}
	log.Println("Generator workers started.")
	numRhsWorkers := int(*NumberOfWorkers)
	var wg sync.WaitGroup
	for lhs := range queue {
		for i := 0; i < numRhsWorkers; i++ {
			wg.Add(1)
			go func(id int) {
				rhsWorkerCount++
				defer func() {
					rhsWorkerCount--
					completedPasses++
					wg.Done()
				}()
				rhs, _ := counters.NewByteCounter(int(*keySpaceSize))
				for func() { _ = rhs.Set(0, byte(id)) }(); true; func() { _ = rhs.Add(numRhsWorkers) }() {
					if lhs.raw == rhs.String() {
						return
					}
					if rhsHash := rhs.Sha1(); lhs.hash == rhsHash {
						log.Printf("collision\n"+
							"lhs %v\n"+
							"rhs %v\n---\n"+
							"%v\n"+
							"---\n"+
							"%v\n"+
							"---\n",
							lhs.hash, rhsHash, lhs.raw, rhs.String())
						os.Exit(1)
					}
					count++
				}
			}(i)
		}
		wg.Wait()
	}
	log.Println("terminating")
}
