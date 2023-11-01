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
	hashScanner "github.com/sam-caldwell/monorepo/go/crsce/tools/find_collisions/lib"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	logInterval         = 1 * time.Second
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

	queue := make(chan struct {
		raw  string
		hash string
	}, 1024)

	for i := uint(0); i < *NumberOfWorkers; i++ {
		go func() {
			c, _ := counters.NewByteCounter(1024)
			for {
				raw := c.String()
				hash := c.Sha1()
				if err := c.Increment(); err != nil {
					return
				}
				queue <- struct {
					raw  string
					hash string
				}{
					raw, hash,
				}
			}
		}()
	}
	log.Println("Generator workers started.")

	func() {
		log.Println("scanner worker starting")
		scanner.Test(queue)
	}()
	log.Printf("All workers terminated.  Remaining Records: %d", queue.Count())
}
