package main

import (
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	defaultKeySpaceSize      = 1024
	initialLogDelay          = 5 * time.Second
	metricsReportingInterval = 10 * time.Second
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
}

type Collector struct {
	startTime int64
	prevCount int64
	metrics   []Metric
}

func (c *Collector) Report(numberWorkers int) {
	var currCount int64
	duration := time.Now().Unix() - c.startTime
	for id := 0; id < numberWorkers; id++ {
		currCount += c.metrics[id].lhsCount
	}
	sample := strings.TrimLeft(c.metrics[0].lhsSample, "0")
	log.Printf("t:%4d, currCount: %12d, prevCount: %12d, chgOps: %12d, sample: %s (%.1f bytes)",
		duration, currCount, c.prevCount, currCount-c.prevCount, sample, float64(len(sample))/2)
	c.prevCount = currCount
}

// AsynchronousJob - Perform a collision search of a given pattern
func AsynchronousJob(id, workerCount, keySpaceSize int, collector *Collector, result chan<- Finding) {
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
	// Set the initial value of LHS to be the worker Id (0-255) at the least-significant byte.
	// This will offset each worker by a value of 1.  Our increment (later) will be id * workerCount
	// So that each worker's search pattern will interleave entire key space.
	//
	if err := lhs.Set(0, byte(id)); err != nil {
		result <- Finding{
			id:        id,
			collision: false,
			err:       err,
		}
		return
	}
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
	for {
		collector.metrics[id].lhsCount++
		//collector.metrics[id].lhsStart = time.Now().Unix()
		collector.metrics[id].lhsSample = lhs.String()
		for {
			//
			// When LHS == RHS, get the next LHS
			//
			if lhs.Equal(&rhs) {
				break
			}
			//
			// When LHS != RHS but the two have the same hash, we have a collision.
			//
			if lhs.Sha1() == rhs.Sha1() {
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
		if err := lhs.Add(workerCount); err != nil {
			result <- Finding{
				id:        id,
				err:       err,
				collision: false,
			}
			break
		}
		//collector.metrics[id].lhsStop = time.Now().Unix()
	} /* LHS Loop */
}
func ReportMetrics(metric *[]Metric) {

}

/*
 * Main routine.
 */
func main() {
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
		time.Sleep(initialLogDelay)
		log.Printf("logging started...")
		metricReportingTimer := time.NewTicker(metricsReportingInterval)
		defer metricReportingTimer.Stop()
		for {
			select {
			case <-metricReportingTimer.C:
				collector.Report(numCpu)
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
		go AsynchronousJob(workerId, numCpu, defaultKeySpaceSize, &collector, results)
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
			log.Printf("collision found (id: %d) (%v) in %v",
				thisResult.id, thisResult.hash, thisResult.raw)
		} else {
			log.Printf("worker %d finished", checkedIn)
			checkedIn++ // We finished, but we have not encountered any collisions or errors.
		}
	}
	log.Printf("terminate with no errors")
}
