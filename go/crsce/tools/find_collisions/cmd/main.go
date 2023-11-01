package main

/*
 * FindCollisions Command
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Create a command to scan the keyspace of a given size
 * across a given number of concurrent workers.
 */

import (
	"bytes"
	"flag"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"runtime"
	"time"
)

const (
	logInterval         = 5 * time.Second
	defaultKeySpaceSize = 1024
	candidateQueueSize  = 65536
)

type Candidate struct {
	raw  []byte
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

	// Timed logging
	go func() {
		log.Printf("Log interval: %d", logInterval)
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				chgCompletedPasses := float64(completedPasses) / elapsedTime

				log.Printf("elapsed: %f objectCnt: %d, object/sec: %6.2f "+
					"queueSz:%d passes/sec: %6.4f",
					elapsedTime, count, float64(count)/elapsedTime, len(queue), chgCompletedPasses)
			}
		}
	}()

	// Generate LHS values to test agains
	go func() {
		const generatorCount = 3
		for worker := 0; worker < generatorCount; worker++ {
			go func(id int) {
				lhsCounter, _ := counters.NewByteCounter(int(*keySpaceSize))
				lhsCounter.Set(0, byte(id))
				for {
					queue <- Candidate{
						raw:  lhsCounter.Bytes(),
						hash: lhsCounter.Sha1(),
					}
					if err := lhsCounter.Add(generatorCount); err != nil {
						log.Println("LHS enumeration complete.")
						return
					}
				}
			}(worker)
		}
	}()

	//Test RHS against LHS values
	exit := make(chan bool, 1)
	for worker := uint(0); worker < *NumberOfWorkers; worker++ {
		log.Printf("Start worker %d", worker)
		for candidate := range queue {
			go func(id uint, lhs *Candidate) {
				rhs, _ := counters.NewByteCounter(int(*keySpaceSize))
				for {
					if bytes.Compare(rhs.Bytes(), lhs.raw) >= 0 {
						completedPasses++
						break
					}
					if lhs.hash == rhs.Sha1() {
						log.Printf("collision\n"+
							"lhs %v\n"+
							"rhs %v\n---\n"+
							"%v\n"+
							"---\n"+
							"%v\n"+
							"---\n",
							lhs.hash, rhs.Sha1(),
							lhs.raw, rhs.String())
						exit <- true
					}
					count++
					_ = rhs.Increment()
				}
			}(worker, &candidate)
		}
	}
	<-exit
}
