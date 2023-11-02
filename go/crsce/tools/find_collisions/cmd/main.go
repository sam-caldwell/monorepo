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
	"encoding/hex"
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"math"
	"os"
	"runtime"
	"time"
)

const (
	logInterval         = 5 * time.Second
	defaultKeySpaceSize = 1024
	candidateQueueSize  = 32768
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

	var count uint64

	var ops float64
	var minOps = math.MaxFloat64
	var maxOps float64

	var passStart int64
	var passStop int64

	var queueSz = candidateQueueSize
	var minQueueSz = math.MaxInt
	var maxQueueSz int

	var workers int
	var minWorkers = 300
	var maxWorkers = int(*NumberOfWorkers) * minWorkers

	var backOffCount int64
	var backOffIncreases int64
	var backOffDecreases int64
	var workerBackoff = time.Millisecond * 1

	var collisionCount int64

	startTime := time.Now().Unix()
	exit := make(chan bool, 1)
	queue := make(chan Candidate, candidateQueueSize)

	// Metrics collection
	go func() {
		time.Sleep(5 * time.Second)
		metricsTicket := time.NewTicker(1 * time.Second)
		defer metricsTicket.Stop()
		for {
			select {
			case <-metricsTicket.C:
				//
				// operations per second (min and max)
				//
				if (ops < minOps) || ((minOps == 0) && (ops < maxOps)) {
					minOps = ops
				}
				if ops > maxOps {
					maxOps = ops
				}
				//
				// queue size stats
				//
				queueSz = len(queue)
				if (queueSz < minQueueSz) || ((minQueueSz == 0) && (queueSz < maxQueueSz)) {
					minQueueSz = queueSz
				}
				if queueSz > maxQueueSz {
					maxQueueSz = queueSz
				}
			}
		}
	}()

	// Timed logging
	go func() {
		log.Printf("Log interval: %d", logInterval)
		time.Sleep(logInterval)
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				passRate := float64(passStop) / elapsedTime
				ops = float64(count) / elapsedTime

				workers = runtime.NumGoroutine()
				log.Printf("{"+
					"\"t\": %4.f, "+
					"\"obj\":{"+
					"\"cnt\": %8d, \"rate/sec\": %8.f, \"min\": %8.f, \"max\": %8.f"+
					"}, "+
					"\"queue\":{"+
					"\"sz\":%5d, \"min\": %5d, \"max\": %5d"+
					"}, "+
					"\"pass\":{"+
					"\"start\":%5d, \"stops\":%5d, \"rate/s\": %4.2f"+
					"}, "+
					"\"worker\":{"+
					"\"count\": %5d \"max\": %5d"+
					"}, "+
					"\"backoff\":{"+
					"\"count\": %5d, \"backoff\": %8v, \"inc\": %5d, \"dec\": %5d"+
					"},"+
					"\"collisionCount\": %3d"+
					"}",
					elapsedTime,
					count, ops, minOps, maxOps,
					queueSz, minQueueSz, maxQueueSz,
					passStart, passStop, passRate,
					workers, maxWorkers,
					backOffCount, workerBackoff, backOffIncreases, backOffDecreases,
					collisionCount)
			}
		}
	}()

	// Generate LHS values to test against
	go func() {
		const generatorCount = 4
		for worker := 0; worker < generatorCount; worker++ {
			go func(id int) {
				lhsCounter, _ := counters.NewByteCounter(int(*keySpaceSize))
				_ = lhsCounter.Set(0, byte(id))
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
			time.Sleep(1 * time.Second)
		}
	}()

	//Test RHS against LHS values
	log.Printf("Start consumers")
	for candidate := range queue {
		passStart++
		go func(lhs Candidate) {
			rhs, _ := counters.NewByteCounter(int(*keySpaceSize))
			for {
				if bytes.Compare(rhs.Bytes(), lhs.raw) >= 0 {
					passStop++
					break
				}
				if lhs.hash == rhs.Sha1() {
					collisionCount++
					collisionFound(lhs.hash, rhs.Sha1(), lhs.raw, rhs.Bytes())
				}
				count++
				_ = rhs.Increment()
			}
		}(candidate)

		//BackOff if needed.
		for runtime.NumGoroutine() >= maxWorkers {
			backOffCount++
			time.Sleep(workerBackoff)
			//
			// Adjust backoff
			//
			if runtime.NumGoroutine() > 9*maxWorkers/10 {
				ansi.Red()
				workerBackoff *= 2
				backOffIncreases++
			} else {
				ansi.Green()
				if runtime.NumGoroutine() < minWorkers {
					workerBackoff = 1
				} else {
					if workerBackoff > 1 {
						workerBackoff /= 2
					}
				}
				backOffDecreases++
			}
		}
	}
	//}
	<-exit
}

func collisionFound(lhsHash, rhsHash string, lhsRaw, rhsRaw []byte) {
	log.Printf("collision\n"+
		"lhs %v\n"+
		"rhs %v\n"+
		"---\n"+
		"%v\n"+
		"---\n"+
		"%v\n"+
		"---\n",
		lhsHash, rhsHash,
		hex.EncodeToString(lhsRaw), hex.EncodeToString(rhsRaw))
	os.Exit(1)
}
