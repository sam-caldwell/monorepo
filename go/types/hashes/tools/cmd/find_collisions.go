package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	defaultKeySpaceSize      = 1024
	initialLogDelay          = 5 * time.Second
	timeWindow               = 10
	metricsReportingInterval = timeWindow * time.Second
)

// Finding - A finding is a report object describing some outcome of the AsynchronousJob()
type Finding struct {
	id        int
	collision bool
	err       error
	hash      string
	raw       string
}

type Metric struct {
	lhsCount  int64
	lhsSample string
	rhsSample string
}

type Collector struct {
	startTime int64
	prevCount int64
	metrics   []Metric
}

// ParseSeed - parse raw hex-encoded string for byte values representing a seed number
func ParseSeed(raw *string) ([]byte, error) {
	if *raw == "" {
		log.Fatal("a seed must be specified")
	}
	decoded, err := hex.DecodeString(*raw)
	if err != nil {
		log.Fatalf("decoding error: %v", err)
	}
	length := len(decoded)
	reversed := make([]byte, length)
	for i := 0; i < length; i++ {
		reversed[length-i-1] = decoded[i]
	}
	return reversed, err
}

// AsynchronousJob - Perform a collision search of a given pattern
func AsynchronousJob(segment, id, workerCount, segmentCount, keySpaceSize int, seed []byte, collector *Collector, result chan<- Finding) {
	//
	// Create the LHS (Left-hand side) counter.  This will be the counter we
	// compare against.
	//
	lhs, err := counters.NewByteCounter(keySpaceSize)
	if err != nil {
		result <- Finding{
			id:        id,
			collision: false,
			err:       err,
		}
		return
	}
	//
	// Set the lhs to the seed value, which will be over-ridden by the workerId.
	//
	if err := lhs.SetBytes(0, seed); err != nil {
		result <- Finding{
			id:        id,
			collision: false,
			err:       err,
		}
		return
	}
	//
	// Set the initial value of LHS to be the worker id (0-255) at the least-significant byte.
	// This will offset each worker by a value of 1.  Our increment (later) will be id * workerCount
	// So that each worker's search pattern will interleave entire key space.
	//
	if err := lhs.Add(id + segment); err != nil {
		result <- Finding{
			id:        id,
			collision: false,
			err:       err,
		}
		return
	}
	log.Printf("lhsStart: %v", strings.TrimLeft(lhs.String(), "0"))
	//
	// Create the Right-Hand Side counter.  This will be the counter we will use to
	// iterate over {0,...,LHS}.  Thus, for every LHS, we will iterate over every value from 0 to LHS
	// testing for a collision.
	//
	rhs, err := counters.NewByteCounter(keySpaceSize)
	if err != nil {
		result <- Finding{
			id:        id,
			collision: false,
			err:       err,
		}
		return
	}
	//
	// Loop through all values of LHS starting with the initial value.
	// When lhs.Add() is executed, the increment will be the product of the worker id
	// and the number of workers running.
	//
	var a [20]byte
	var b [20]byte
	for {
		collector.metrics[id].lhsCount++
		collector.metrics[id].lhsSample = lhs.String()
		for {
			collector.metrics[id].rhsSample = rhs.String()
			//
			// When LHS == RHS, get the next LHS
			//
			if lhs.Equal(&rhs) {
				break
			}
			//
			// When LHS != RHS but the two have the same hash, we have a collision.
			//
			a = lhs.Sha1Bytes()
			b = rhs.Sha1Bytes()
			if bytes.Equal(a[:], b[:]) {
				result <- Finding{
					id:        id,
					hash:      lhs.Sha1(),
					raw:       lhs.String(),
					collision: true,
				}
				return
			}
			//
			// Increment RHS to the next value and carry on until exhausted.
			//
			if err := rhs.Increment(); err != nil {
				break //All RHS values have been exhausted.
			}
		} /* RHS Loop */
		//
		//LHS.Add() increments the LHS value by id*workerCount and when an error occurs, we are exhausted.
		//
		if err := lhs.Add(workerCount * segmentCount); err != nil {
			result <- Finding{
				id:        id,
				err:       err,
				collision: false,
			}
			break
		}
	} /* LHS Loop */
}

// main - the main routine for a single unit of processing.
func main() {
	rawSeed := flag.String("Seed", "", "Seed value (hex-encoded string)")
	startingWorkerId := flag.Int("Segment", 0, "The segment starting worker Id")
	segmentCount := flag.Int("SegmentCount", 0, "This is the planned segment count")
	flag.Parse()
	seed, err := ParseSeed(rawSeed)
	if err != nil {
		log.Fatalf("Seed Parse error: %v", err)
	}
	log.Printf("Seed: %v (%d)", hex.EncodeToString(seed), len(seed))
	log.Println("Initializing...")
	//
	// The number of CPUs will determine the number
	// of workers we will distribute the problem across.
	//
	numCpu := runtime.NumCPU()
	log.Printf("Number CPU: %d", numCpu)
	if runtime.NumCPU() > 255 {
		panic("more than 255 CPUs not supported at this time")
	}
	results := make(chan Finding, numCpu)
	runtime.GOMAXPROCS(numCpu)
	//
	// Create an array of Metric objects (one per worker).
	collector := Collector{
		metrics: make([]Metric, numCpu),
	}
	collector.startTime = time.Now().Unix()
	//
	// Log the metrics via a goroutine using a timer
	//
	go func() {
		var currCount int64
		//
		// ToDo: setup file logging so we can centralize things for analysis...
		//
		log.Printf("logging started...")
		metricReportingTimer := time.NewTicker(metricsReportingInterval)
		defer metricReportingTimer.Stop()
		for {
			select {
			case <-metricReportingTimer.C:
				rand.Seed(time.Now().UnixNano())
				sampleId := rand.Intn(numCpu)
				duration := time.Now().Unix() - collector.startTime
				for id := 0; id < numCpu; id++ {
					currCount += collector.metrics[id].lhsCount
				}
				lhsSample := strings.TrimLeft(collector.metrics[sampleId].lhsSample[5:], "0")
				rhsSample := strings.TrimLeft(collector.metrics[sampleId].rhsSample[5:], "0")
				log.Printf("t:%4d, currCount: %12d, prevCount: %12d, chg Ops: %12.f, "+
					"id: %2d, lhs:%d:%s (%.f bytes) rhs: %s (%.f bytes)",
					duration,
					currCount,
					collector.prevCount,
					float64(currCount-collector.prevCount)/float64(timeWindow),
					sampleId,
					*startingWorkerId,
					lhsSample,
					math.Ceil(float64(len(lhsSample))/2),
					rhsSample,
					math.Ceil(float64(len(rhsSample))/2))
				collector.prevCount = currCount
				currCount = 0
			}
		}
	}()
	//
	// Loop over the number of CPUs dispatching workers accordingly.
	// We trust that the linux kernel and golang runtime will ensure
	// each worker is mapped to a single core.
	//
	for workerId := 0; workerId < numCpu; workerId++ {
		log.Printf("Start worker %d", workerId)
		go AsynchronousJob(*startingWorkerId, workerId, *segmentCount, numCpu, defaultKeySpaceSize, seed, &collector, results)
	}
	//
	// Block until all results are in or a single error is encountered.
	// Each worker is allowed to report one result.
	//
	log.Printf("Waiting on results")
	for checkedIn := 0; checkedIn < numCpu; {
		log.Printf("\tWaiting...")
		thisResult := <-results
		if thisResult.err != nil {
			log.Printf("Error(id: %d): %v", thisResult.id, thisResult.err)
			os.Exit(1)
		}
		if thisResult.collision {
			//report the collision and keep searching.
			log.Printf("collision found (id: %2d) (%v) in %v",
				thisResult.id, thisResult.hash, thisResult.raw)
		} else {
			log.Printf("worker %d finished", checkedIn)
			checkedIn++ // We finished, but we have not encountered any collisions or errors.
		}
	}
	log.Printf("terminate with no errors")
}
