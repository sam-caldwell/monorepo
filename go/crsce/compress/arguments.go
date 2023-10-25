package main

/*
 * CRSCE Compress Tool
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Proof of concept for CRSCE compression.
 */

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

type Arguments struct {
	In        *string
	Out       *string
	BlockSize *uint
}

func (args *Arguments) GetArgs() {
	const (
		MinSize = 1024
		MaxSize = 65535
	)
	shortHelp := flag.Bool("-h", false, "Display Usage")
	longHelp := flag.Bool("--help", false, "Display Usage")
	args.In = flag.String("--in", words.EmptyString, "input file")
	args.Out = flag.String("--out", words.EmptyString, "output file")
	args.BlockSize = flag.Uint("--bs", 1024, "")

	flag.Parse()

	exit.OnCondition(*shortHelp || *longHelp, exit.Success, "usage:", commandUsage)

	exit.OnCondition(strings.TrimSpace(*args.In) == "", exit.InvalidInput, "input file cannot be empty", commandUsage)
	exit.OnCondition(!file.Exists(*args.In), exit.InvalidInput, "input file does not exist.", commandUsage)

	exit.OnCondition(strings.TrimSpace(*args.Out) == "", exit.InvalidInput, "output file cannot be empty", commandUsage)
	exit.OnCondition(file.Exists(*args.Out), exit.InvalidInput, "output file should not exist.", commandUsage)

	exit.OnCondition(*args.BlockSize <= MinSize, exit.InvalidInput, fmt.Sprintf("minimum block size %d", MinSize), commandUsage)
	exit.OnCondition(*args.BlockSize > MaxSize, exit.InvalidInput, fmt.Sprintf("minimum block size %d", MinSize), commandUsage)
}
