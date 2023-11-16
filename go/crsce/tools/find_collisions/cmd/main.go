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
	"github.com/sam-caldwell/monorepo/go/exit"
	"log"
	"math"
	"os"
	"runtime"
	"strings"
	"time"
)

import "github.com/newrelic/go-agent/v3/newrelic"

const (
	initialBackoff      = 1 * time.Millisecond
	maxBackoff          = 4 * time.Minute
	generatorCount      = 16
	logInterval         = 20 * time.Second
	defaultKeySpaceSize = 1024
	candidateQueueSize  = 1048576
)

type Candidate struct {
	hash string
	raw  []byte
}

func main() {
	app, err := newrelic.NewApplication(
  		newrelic.ConfigAppName("find_collisions"),
  		newrelic.ConfigLicense("<redacted>"),
  		newrelic.ConfigAppLogForwardingEnabled(true),
	)
        if err!= nil {
		panic(err)
	}


	keySpaceSize := flag.Uint(
		"keySpaceSize",
		defaultKeySpaceSize,
		"Number of bytes in the key space to scan")

	NumberOfWorkers := flag.Uint(
		"numberWorkers",
		uint(runtime.NumCPU()),
		"Number of workers to launch")

	flag.Parse()

	runtime.GOMAXPROCS(80)

	if *NumberOfWorkers > 255 {
		log.Println("number of workers exceeds max worker count")
		os.Exit(exit.GeneralError)
	}

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
	var minWorkers = 400
	var maxWorkers = int(*NumberOfWorkers) * minWorkers
	var backOffTime time.Duration
	var backOffCount int64
	var backOffIncreases int64
	var backOffDecreases int64
	var workerBackoff = initialBackoff
	var collisionCount int64
	var currentCandidate []byte
	startTime := time.Now().Unix()
	done := make(chan bool, 1)
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
		log.Printf("Log interval: %v", logInterval)
		time.Sleep(logInterval)
		ticker := time.NewTicker(logInterval) // Adjust the interval as needed
		defer ticker.Stop()
		totalBackOffTime := float64(0)
		for {
			select {
			case <-ticker.C:
				elapsedTime := float64(time.Now().Unix() - startTime)
				passRate := float64(passStop) / elapsedTime
				ops = float64(count) / elapsedTime

				totalBackOffTime += float64(backOffTime)

				workers = runtime.NumGoroutine()
				log.Printf("{"+
					"\"t\": %4.f, "+
					"\"obj\":{"+
					"\"cnt\": %12d, \"rate/sec\": %12.f, \"min\": %12.f, \"max\": %12.f"+
					"}, "+
					"\"queue\":{"+
					"\"sz\":%5d, \"min\": %5d, \"max\": %5d"+
					"}, "+
					"\"pass\":{"+
					"\"start\":%5d, \"stops\":%5d, \"rate/s\": %06.2f"+
					"}, "+
					"\"worker\":{"+
					"\"count\": %5d \"max\": %5d"+
					"}, "+
					"\"backoff\":{"+
					"\"count\": %5d, \"backoff\": %8v, \"inc\": %5d, \"dec\": %5d, \"time\":%10v, "+
					"}, "+
					"\"collisionCount\": %3d, "+
					"\"c\":%s"+
					"}",
					elapsedTime,
					count, ops, minOps, maxOps,
					queueSz, minQueueSz, maxQueueSz,
					passStart, passStop, passRate,
					workers, maxWorkers,
					backOffCount, workerBackoff, backOffIncreases, backOffDecreases, backOffTime.Truncate(time.Microsecond),
					collisionCount,
					strings.TrimRight(strings.TrimLeft(hex.EncodeToString(currentCandidate), "0"), "0"))
				ansi.Reset()
			}
		}
	}()

	// Generate LHS values to test against
	go func() {
		lhsCounter, _ := counters.NewByteCounter(int(*keySpaceSize))
		for {
			txnLhs := app.StartTransaction("lhs_generation")
			queue <- Candidate{
				raw:  lhsCounter.Bytes(),
				hash: lhsCounter.Sha1(),
			}
			if err := lhsCounter.FastIncrement(); err != nil {
				break
			}
			txnLhs.End()
		}
	}()
	//Give the LHS time to populate the queue
        time.Sleep(5*time.Second)


	//Test RHS against LHS values
	log.Printf("Start consumers")

	for candidate := range queue {
		txnRhs := app.StartTransaction("rhs_consumers")

		passStart++
		go func(lhs Candidate) {
			rhs, _ := counters.NewByteCounter(int(*keySpaceSize))
			for {
				if bytes.Compare(lhs.raw, rhs.Bytes()) <= 0 {
					passStop++
					currentCandidate = lhs.raw
					if runtime.NumGoroutine() < minWorkers {
						backOffDecreases++
						workerBackoff = initialBackoff
					}
					break
				}
				if lhs.hash == rhs.Sha1() {
					collisionCount++
					collisionFound(lhs.hash, rhs.Sha1(), lhs.raw, rhs.Bytes())
				}
				if err := rhs.Increment(); err != nil {
					break
				}
				count++
			}
		}(candidate)

		//BackOff if needed.
		backOffStart := time.Now()
		for runtime.NumGoroutine() >= maxWorkers {
			backOffCount++
			if workerBackoff > maxBackoff {
				ansi.Green()
				backOffDecreases++
				workerBackoff /= 2
			}
			time.Sleep(workerBackoff)
			if runtime.NumGoroutine() > maxWorkers {
				ansi.Red()
				workerBackoff *= 2
				backOffIncreases++
				continue
			}
			ansi.Reset()
		}
		backOffTime = time.Since(backOffStart)
	       	txnRhs.End()
	}
	//}
	<-done
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
	os.Exit(exit.GeneralError)
}
