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
	"math/big"
	"os"
	"path/filepath"
)

const (
	root         = "/tmp/data"
	sha1file     = "sha1"
	keySpaceSize = 1024
)

func generateSHA1(input *big.Int) (dirPath string) {
	hash := sha1.Sum(input.Bytes())
	h := hex.EncodeToString(hash[:])

	// Divide the hash into directories and filename
	for i := 0; i < len(h); i += 4 {
		dirPath = filepath.Join(dirPath, h[i:i+4])
	}
	return dirPath
}

func main() {
	log.Println("Starting")
	// Create a directory to store the generated hashes
	if directory.Exists(root) {
		log.Println("Cleaning")
		_ = os.RemoveAll(root)
	}
	_ = os.MkdirAll(root, 0744)
	// Iterate through strings of length 0 to 1024
	maxValue := new(big.Int)
	maxValue.Exp(big.NewInt(2), big.NewInt(8*keySpaceSize), nil) // 2^(8*keySpaceSize)
	for i := new(big.Int); i.Cmp(maxValue) < 0; i.Add(i, big.NewInt(1)) {
		// Generate SHA-1 hash
		dirPath := generateSHA1(i)
		//log.Printf("i: %d (%s)", i, dirPath)
		// Create directories
		if err := os.MkdirAll(filepath.Join(root, dirPath), 0744); err != nil {
			panic(err)
		}

		func() {
			// Create a file with the SHA-1 hash as the name
			fileName := filepath.Join(root, dirPath, sha1file)
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
