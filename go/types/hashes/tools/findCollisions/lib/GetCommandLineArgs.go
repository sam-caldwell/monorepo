package findCollision

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/version"
	"log"
	"os"
)

func GetCommandLineArgs() ([]byte, int, int) {

	var err error
	var seed []byte

	versionFlag := flag.Bool("version", false, "Show the current version")

	rawSeed := flag.String("Seed", "", "Seed value (hex-encoded string)")

	startingWorkerId := flag.Int("Segment", 0, "The segment starting worker Id")

	segmentCount := flag.Int("SegmentCount", 0, "This is the planned segment count")

	flag.Parse()
	if *versionFlag {
		fmt.Println(version.Version)
		os.Exit(exit.Success)
	}

	if seed, err = ParseSeed(rawSeed); err != nil {
		log.Fatalf("Seed Parse error: %v", err)
	}
	log.Printf("Seed: %v (%d)", hex.EncodeToString(seed), len(seed))
	return seed, *startingWorkerId, *segmentCount

}
