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
	"time"
)

const (
	logInterval         = 1 * time.Second
	defaultKeySpaceSize = 1024
	candidateQueueSize  = 128
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
	queue := make(chan Candidate, 1024)

	go func() {
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				log.Printf("elapsed: %f objectCnt: %d, object/sec: %6.2f queueSz:%d ",
					elapsedTime, count, float64(count)/elapsedTime, len(queue))
			}
		}
	}()

	for i := uint(0); i < *NumberOfWorkers; i++ {
		go func(offset uint) {
			c, _ := counters.NewByteCounter(1024)
			_ = c.Set(0, byte(offset))
			for {
				queue <- Candidate{
					raw:  c.String(),
					hash: c.Sha1(),
				}
				if err := c.Add(int(*NumberOfWorkers)); err != nil {
					return
				}
			}
		}(i)
		log.Printf("worker %d started", i)
	}
	log.Println("Generator workers started.")
	for lhs := range queue {
		for i := uint(0); i < *NumberOfWorkers; i++ {
			go func() {
				rhs, _ := counters.NewByteCounter(1024)
				for func() { _ = rhs.Set(0, byte(i)) }(); lhs.raw != rhs.String(); func() { _ = rhs.Add(int(*NumberOfWorkers)) }() {
					if lhs.hash == rhs.Sha1() {
						if lhs.raw == rhs.String() {
							log.Println("Pass complete")
							return
						}
						log.Printf("collision\n"+
							"lhs %v\n"+
							"rhs %v\n\n---\n%v\n---\n\n---\n%v\n---\n",
							lhs.hash, rhs.Sha1(), lhs.raw, rhs.String())
						os.Exit(1)
					}
					count++
				}
			}()
		}
	}
	log.Println("terminating")
}
