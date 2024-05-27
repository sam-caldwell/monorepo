package crsce

/*
 * Argument GetArgs() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Parse the command-line arguments and validate the inputs.
 */

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// GetArgs - Parse cli arguments and validate the values.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (args *Arguments) GetArgs() {
	const (
		MinSize = 1048576 //Bytes (1MB) ( 8,388,608 bits) sqrt(1MB) => 1024
		MaxSize = 8388608 //Bytes (8MB) (67,108,864 bits) sqrt(8MB) => 2896.3 => 2897
	)
	shortHelp := flag.Bool("h", false, "Display Usage")
	longHelp := flag.Bool("help", false, "Display Usage")
	inFile := flag.String("in", words.EmptyString, "input file")
	outFile := flag.String("out", words.EmptyString, "output file")
	blockSize := flag.Uint("bs", defaultBlockSize, "")
	flag.Parse()

	exit.OnCondition(*shortHelp || *longHelp, exit.Success, "usage:", commandUsage)

	args.In = *inFile
	exit.OnCondition(strings.TrimSpace(args.In) == "", exit.InvalidInput, "input file cannot be empty", commandUsage)
	exit.OnCondition(!file.Exists(args.In), exit.InvalidInput, "input file does not exist.", commandUsage)

	args.Out = *outFile
	exit.OnCondition(strings.TrimSpace(args.Out) == "", exit.InvalidInput, "output file cannot be empty", commandUsage)
	exit.OnCondition(file.Exists(args.Out), exit.InvalidInput, "output file should not exist.", commandUsage)

	args.BlockSize = *blockSize
	exit.OnCondition(args.BlockSize < MinSize, exit.InvalidInput,
		fmt.Sprintf("minimum block size %d (got %d)", MinSize, *blockSize), commandUsage)
	exit.OnCondition(args.BlockSize > MaxSize, exit.InvalidInput,
		fmt.Sprintf("maximum block size %d (got %d)", MinSize, *blockSize), commandUsage)
	exit.OnCondition((args.BlockSize%8) != 0, exit.InvalidInput,
		fmt.Sprintf("block size (%d) must be evenly divisible by 8 bits", args.BlockSize), commandUsage)
}
