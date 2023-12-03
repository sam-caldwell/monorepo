package main

import (
	"bytes"
	"flag"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types/hashes"
	"log"
	"os"
)

const (
	cmdUsage = ""
)

func main() {
	hashFile := flag.String("file", "", "the hash file to load into memory")
	searchFor := flag.String("searchFor", "", "the hash to search for in the table")
	flag.Parse()
	exit.OnCondition(*hashFile == "", exit.InvalidInput, "hash file must be defined", cmdUsage)
	exit.OnCondition(!file.Existsp(hashFile), exit.NotFound, "file not found", cmdUsage)

	exit.OnCondition(*searchFor == "", exit.InvalidInput, "provide a sha-1 hash to search for", cmdUsage)
	exit.OnCondition(len(*searchFor) != 40, exit.InvalidInput, "hash must be 20-bytes hex-encoded", cmdUsage)

	log.Printf("loading file (%s)...", *hashFile)

	size := file.GetFileSizeByName(*hashFile)
	recordCount, err := convert.Int64ToIntSafe(size / 20) // 20-bytes per hash
	exit.OnError(err, exit.GeneralError, "Invalid record count")

	table := make([]hashes.Sha1, recordCount)
	log.Printf("Table Size:%d", len(table))

	if binarySearch(&table, searchFor) {
		log.Println("FOUND")
		os.Exit(exit.Success)
	} else {
		log.Println("NOT FOUND")
		os.Exit(exit.NotFound)
	}
}

func binarySearch(table *[]hashes.Sha1, searchFor *string) (ok bool) {
	var criteria hashes.Sha1
	criteria.HexEncodedString(searchFor)

	endPos := len(*table)
	pos := endPos / 2
	for {
		switch bytes.Compare((*table)[pos][:], criteria[:]) {
		case -1:
			// we have reached the beginning
			if pos <= 0 {
				return false
			} else {
				//search left
				pos = pos / 2
			}
		case 0:
			return true
		case 1:
			// we have reached the end.
			if pos >= endPos {
				return false
			} else {
				// search right
				pos += pos / 2
			}
		}
	}
}
