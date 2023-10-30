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
	root          = "/opt/data/crsce"
	fileExtension = ".sha1"
	keySpaceSize  = 1024
)

func initialize() {
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
	//var fileCount uint64

	log.Println("Starting")
	initialize()
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

			hashStart := time.Now().UnixNano()
			hash := input.Sha1()

			dirPath := filepath.Join(
				root,
				hash[0:2], hash[2:4], hash[4:6], hash[6:8], hash[8:10],
				hash[10:12], hash[12:14], hash[14:16], hash[16:18], hash[18:20])
			hashElapsed := time.Now().UnixNano() - hashStart

			diskOpStart := time.Now().UnixNano()
			if directory.Existsp(&dirPath) {
				log.Fatalf("collision found at (%s): %v", dirPath, err)
			}
			if err := os.MkdirAll(dirPath, 0744); err != nil {
				log.Fatalf("Failed to create path (%s): %v", dirPath, err)
			}
			diskOpElapsed := time.Now().UnixNano() - diskOpStart

			directoryCount++
			stopTime := time.Now().Unix()

			go func() {
				log.Printf("objects: %d, object/sec: %5.2f hashTime:%d diskOpElapsed:%d",
					directoryCount,
					float64(directoryCount)/float64(stopTime-startTime), hashElapsed, diskOpElapsed)
			}()
		}()
	}
}
