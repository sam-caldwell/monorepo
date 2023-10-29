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
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"log"
	"os"
	"path/filepath"
)

const (
	root          = "/opt/data/crsce"
	fileExtension = ".sha1"
	keySpaceSize  = 1024

	partitionSize = 10
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
	log.Println("Starting")
	initialize()
	// Iterate through strings of length 0 to 1024
	input, err := counters.NewByteCounter(keySpaceSize)
	if err != nil {
		panic(err)
	}
	for {
		if err := input.Increment(); err != nil {
			break
		}
		func() {
			hash := input.Sha1()

			dirPath := filepath.Join(root, hash[0:partitionSize])
			fileName := filepath.Join(dirPath, hash[partitionSize:]+fileExtension)

			if err := os.MkdirAll(dirPath, 0744); err != nil {
				log.Fatalf("Failed to create path (%s)", dirPath)
			}
			if file.Exists(fileName) {
				log.Fatalf("Collision found at (%s): %s", fileName, hash)
			}
			fileHandle, err := os.Create(fileName)
			if err != nil {
				log.Fatalf("Error creating file(%s):%s", fileName, err)
				return
			}
			defer func() {
				if err := fileHandle.Close(); err != nil {
					log.Fatalf("error closing file at %s: %v", hash, err)
				}
			}()
		}()
	}
}
