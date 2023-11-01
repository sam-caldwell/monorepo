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
	"fmt"
	hashScanner "github.com/sam-caldwell/monorepo/go/crsce/tools/find_collisions/lib"
	"log"
	"os"
	"runtime"
	"sync"
)

const (
	defaultKeySpaceSize = 1024
	candidateQueueSize  = 128
)

func main() {
	defer func() { _ = os.Stdout.Sync() }()
	var wg sync.WaitGroup
	keySpaceSize := flag.Uint(
		"keySpaceSize",
		defaultKeySpaceSize,
		"Number of bytes in the key space to scan")

	NumberOfWorkers := flag.Uint(
		"numberWorkers",
		uint(runtime.NumCPU()),
		"Number of workers to launch")

	flag.Parse()

	log.Printf("Starting with %d generator workers (keySpaceSz:%d)\n", *NumberOfWorkers, *keySpaceSize)

	var scanner hashScanner.Worker
	if err := scanner.Initialize(0, 0, *keySpaceSize); err != nil {
		log.Println("Failed to initialize scanner")
	}
	log.Println("scanner worker is initialized.")

	workers := make([]hashScanner.Worker, *NumberOfWorkers)

	queue := hashScanner.NewCandidateQueue(candidateQueueSize)

	for i := uint(0); i < *NumberOfWorkers; i++ {
		if err := workers[i].Initialize(i, *NumberOfWorkers, *keySpaceSize); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	}
	log.Println("Generator workers initialized.")
	for i := uint(0); i < *NumberOfWorkers; i++ {
		go workers[i].EnumerateKeyspace(&wg, queue)
		log.Printf("Generator worker (%v) started\n", i)
	}
	log.Println("Generator workers started.")
	log.Println("scanner worker starting")
	scanner.Test(&wg, queue)
	log.Println("Waiting on all workers to terminate")
	wg.Wait()
	log.Printf("All workers terminated.  Remaining Records: %d", queue.Count())
}
