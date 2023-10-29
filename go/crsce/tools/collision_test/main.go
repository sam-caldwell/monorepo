package main

/*
 *
 */

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"log"
	"os"
	"path/filepath"
)

const (
	root         = "data"
	extension    = ".sha1"
	keySpaceSize = 1024
)

func generateSHA1(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func partitionSHA1(hash string) string {
	// Divide the hash into directories and filename
	dirPath := ""
	for i := 0; i < len(hash); i += 4 {
		dirPath += hash[i:i+4] + "/"
	}
	return dirPath + "sha1"
}

func main() {
	// Create a directory to store the generated hashes
	if directory.Exists(root) {
		_ = os.RemoveAll(root)
	}
	_ = os.MkdirAll(root, os.ModeDir)

	// Iterate through strings of length 0 to 1024
	for i := 0; i <= keySpaceSize; i++ {
		var input string
		// Generate SHA-1 hash
		hash := generateSHA1(input)

		// Create directory path based on hash
		dirPath := partitionSHA1(hash)

		// Create directories
		_ = os.MkdirAll(filepath.Join(root, dirPath), os.ModeDir)

		func() {
			// Create a file with the SHA-1 hash as the name
			fileName := filepath.Join(root, dirPath+extension)
			if file.Exists(fileName) {
				log.Fatalf("Collision found at %s", fileName)
			}
			file, err := os.Create(fileName)
			if err != nil {
				log.Fatalf("Error creating file(%s):%s", fileName, err)
				return
			}
			defer func() { _ = file.Close() }()
		}()
	}
}
