package main

/*
 * collision_test
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * Test 1024-bit collisions.
 */

import (
	"github.com/sam-caldwell/monorepo/go/counters"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	root         = "/opt/data/crsce"
	keySpaceSize = 1024
)

func init() {
	if directory.Exists(root) {
		log.Println("Cleaning")
		if err := os.RemoveAll(root); err != nil {
			panic(err)
		}
	}
	// Create a directory to store the generated hashes
	if err := os.MkdirAll(root, 0744); err != nil {
		panic(err)
	}
	log.Printf("Created root directory: %s", root)
}

func main() {
	var directoryCount uint64

	log.Println("Starting")

	// Iterate through strings of length 0 to 1024
	input, err := counters.NewByteCounter(keySpaceSize)
	if err != nil {
		panic(err)
	}
	startTime := time.Now().Unix()
	for {
		if err := input.Increment(); err != nil {
			break
		}
		func() {
			hash := input.Sha1()

			dirPath := filepath.Join(
				root,
				hash[0:2], hash[2:4], hash[4:6], hash[6:8], hash[8:10],
				hash[10:12], hash[12:14], hash[14:16], hash[16:18], hash[18:20])

			if directory.Existsp(&dirPath) {
				log.Fatalf("collision found at (%s): %v", dirPath, err)
			}
			if err := os.MkdirAll(dirPath, 0744); err != nil {
				log.Fatalf("Failed to create path (%s): %v", dirPath, err)
			}

			stopTime := time.Now().Unix()
			directoryCount++

			go func() {

				ticker := time.NewTicker(5 * time.Second) // Adjust the interval as needed
				defer ticker.Stop()

				for {
					select {
					case <-ticker.C:
						log.Printf("objects: %11d, object/sec: %11.2f",
							directoryCount, float64(directoryCount)/float64(stopTime-startTime))
					}
				}
			}()
		}()
	}
}
