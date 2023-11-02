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

	startTime := time.Now().Unix()
	count := uint64(0)
	queue := make(chan Candidate, candidateQueueSize)
	completedPasses := 0
	ops := float64(0)
	minOps := math.MaxFloat64
	maxOps := float64(0)
	queueSz := candidateQueueSize
	minQueueSz := math.MaxInt
	maxQueueSz := 0

	// Metrics collection
	go func() {
		time.Sleep(5 * time.Second)
		metricsTicket := time.NewTicker(1 * time.Second)
		defer metricsTicket.Stop()
		for {
			select {
			case <-metricsTicket.C:
				if (ops < minOps) || ((minOps == 0) && (ops < maxOps)) {
					minOps = ops
				}
				if ops > maxOps {
					maxOps = ops
				}
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
				chgCompletedPasses := float64(completedPasses) / elapsedTime
				ops = float64(count) / elapsedTime
				log.Printf("elapsed: %f objectCnt: %d,"+
					" object/sec: %6.2f (min: %6.2f, max: %6.2f) "+
					"queueSz:%d (min: %d, max: %d)"+
					"passes/sec: %6.4f",
					elapsedTime, count, ops, minOps, maxOps, queueSz,
					minQueueSz, maxQueueSz, chgCompletedPasses)
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
	exit := make(chan bool, 1)
	maxWorkers := int(*NumberOfWorkers) * 200
	workerBackoff := time.Duration(1)
	log.Printf("Start consumers")
	backOffId := int64(0)
	backOffIncreases := int64(0)
	backOffDecreases := int64(0)
	for candidate := range queue {
		//log.Printf("Start cycle. workers (count: %d)", runtime.NumGoroutine())
		go func(lhs Candidate) {
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
						hex.EncodeToString(lhs.raw), rhs.String())
					os.Exit(1)
				}
				count++
				_ = rhs.Increment()
			}
		}(candidate)
		for runtime.NumGoroutine() >= maxWorkers {
			log.Printf("%d worker count: %d (max %d, backoff: %v)",
				backOffId, runtime.NumGoroutine(), maxWorkers, workerBackoff*time.Millisecond)
			time.Sleep(time.Millisecond * workerBackoff)
			if runtime.NumGoroutine() > maxWorkers/2 {
				ansi.Red()
				workerBackoff *= 2
				backOffIncreases++
			} else {
				ansi.Green()
				workerBackoff /= 2
				backOffDecreases++
			}
			log.Printf("%d backoff complete (count: %d inc: %d, dec: %d)",
				backOffId, runtime.NumGoroutine(), backOffIncreases, backOffDecreases)
			ansi.Reset()
			backOffId++
		}
	}
	//}
	<-exit
}
