package crsce

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

const (
	commandUsage = `

crsce-compress --in <input_file> --out <output_file> [--bs <blocksize>]

`
)

const (
	defaultBlockSize = 1048576 //bytes
	s                = 8192    // ceil(sqrt(1048576)) * 8 (bits)
)

// Arguments - Struct representing parsed command-line args.
type Arguments struct {
	In        string
	Out       string
	BlockSize uint
}

// GetArgs - Parse cli arguments and validate the values.
func (args *Arguments) GetArgs() {
	const (
		MinSize = 1024
		MaxSize = 65535
	)
	shortHelp := flag.Bool("-h", false, "Display Usage")
	longHelp := flag.Bool("--help", false, "Display Usage")
	inFile := flag.String("--in", words.EmptyString, "input file")
	outFile := flag.String("--out", words.EmptyString, "output file")
	blockSize := flag.Uint("--bs", defaultBlockSize, "")
	flag.Parse()

	exit.OnCondition(*shortHelp || *longHelp, exit.Success, "usage:", commandUsage)

	args.In = *inFile
	exit.OnCondition(strings.TrimSpace(args.In) == "", exit.InvalidInput, "input file cannot be empty", commandUsage)
	exit.OnCondition(!file.Exists(args.In), exit.InvalidInput, "input file does not exist.", commandUsage)

	args.Out = *outFile
	exit.OnCondition(strings.TrimSpace(args.Out) == "", exit.InvalidInput, "output file cannot be empty", commandUsage)
	exit.OnCondition(file.Exists(args.Out), exit.InvalidInput, "output file should not exist.", commandUsage)

	args.BlockSize = *blockSize
	exit.OnCondition(args.BlockSize <= MinSize, exit.InvalidInput, fmt.Sprintf("minimum block size %d", MinSize), commandUsage)
	exit.OnCondition(args.BlockSize > MaxSize, exit.InvalidInput, fmt.Sprintf("minimum block size %d", MinSize), commandUsage)
}
