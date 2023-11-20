package main

import (
	findCollision "github.com/sam-caldwell/monorepo/go/types/hashes/tools/findCollisions/lib"
	"log"
	"time"
)

const (
	defaultKeySpaceSize = 1024
	//initialLogDelay          = 5 * time.Second
	timeWindow               = 10
	metricsReportingInterval = timeWindow * time.Second
)

// main - the main routine for a single unit of processing.
func main() {
	seed, startingWorkerId, segmentCount := findCollision.GetCommandLineArgs()
	log.Println("Initializing...")
	numberOfWorkers := findCollision.SetNumberCpu()
	results := make(chan findCollision.Finding, numberOfWorkers)
	collector := findCollision.NewCollector(numberOfWorkers)

	go findCollision.TimedLogger(metricsReportingInterval, numberOfWorkers, startingWorkerId, timeWindow, &collector)

	for workerId := 0; workerId < numberOfWorkers; workerId++ {
		go findCollision.AsynchronousJob(startingWorkerId, workerId, segmentCount, numberOfWorkers, defaultKeySpaceSize, seed, &collector, results)
	}

	findCollision.CollectFindings(numberOfWorkers, results)

	log.Printf("terminate with no errors")
}
