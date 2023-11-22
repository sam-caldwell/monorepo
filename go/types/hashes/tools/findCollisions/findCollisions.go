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
	QuickTableSize           = 256 * 256 * 256 * 256 //Table Size: 256 * 256 * 256 * 256 * 20 / 1048576 MB
)

// main - the main routine for a single unit of processing.
func main() {
	seed, startingWorkerId, segmentCount := findCollision.GetCommandLineArgs()
	log.Println("Initializing...")
	numberOfWorkers := findCollision.SetNumberCpu()
	results := make(chan findCollision.Finding, numberOfWorkers)
	collector := findCollision.NewCollector(numberOfWorkers)

	go findCollision.TimedLogger(metricsReportingInterval, numberOfWorkers, startingWorkerId, timeWindow, &collector)

	lookupTable := findCollision.NewQuickTable(defaultKeySpaceSize, QuickTableSize)
	for workerId := 0; workerId < numberOfWorkers; workerId++ {
		log.Printf("Start worker %d", workerId)
		go findCollision.AsynchronousJob(startingWorkerId, workerId, segmentCount, numberOfWorkers,
			defaultKeySpaceSize, seed, lookupTable, &collector, results)
	}

	findCollision.CollectFindings(numberOfWorkers, results)

	log.Printf("terminate with no errors")
}
