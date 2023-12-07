package main

import (
	findCollision2 "github.com/sam-caldwell/monorepo/go/tools/findCollisions/lib"
	"log"
	"time"
)

const (
	defaultKeySpaceSize = 1024
	//initialLogDelay          = 5 * time.Second
	timeWindow               = 10
	metricsReportingInterval = timeWindow * time.Second
	QuickTableSize           = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
)

// main - the main routine for a single unit of processing.
func main() {
	seed, startingWorkerId, segmentCount, hashFileName := findCollision2.GetCommandLineArgs()
	log.Println("Initializing...")
	numberOfWorkers := findCollision2.SetNumberCpu()
	results := make(chan findCollision2.Finding, numberOfWorkers)
	collector := findCollision2.NewCollector(numberOfWorkers)

	lookupTable, rhsStartingSequence := findCollision2.NewQuickTable(hashFileName, defaultKeySpaceSize, QuickTableSize)

	go findCollision2.TimedLogger(metricsReportingInterval, numberOfWorkers, startingWorkerId, timeWindow, &collector)

	for workerId := 0; workerId < numberOfWorkers; workerId++ {
		log.Printf("Start worker %d", workerId)
		go findCollision2.AsynchronousJob(startingWorkerId, workerId, segmentCount, numberOfWorkers,
			defaultKeySpaceSize, seed, lookupTable, rhsStartingSequence, &collector, results)
	}

	findCollision2.CollectFindings(numberOfWorkers, results)

	log.Printf("terminate with no errors")
}
