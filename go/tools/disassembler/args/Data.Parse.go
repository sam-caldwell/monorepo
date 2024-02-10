package args

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
)

// Parse - Parse and validate args then open source and output files.
func (data *Data) Parse(usage string) {

	shortHelp := flag.Bool("-h", false, "show help")
	longHelp := flag.Bool("--help", false, "show help")
	sourceFile := flag.String("--source", words.EmptyString, "Source (binary) file")
	outFile := flag.String("--out", words.EmptyString, "output (disassembled) binary file")
	method := flag.String("--method", words.EmptyString, "Disassembly method (linear|recursive)")
	arch := flag.String("--arch", words.EmptyString, "CPU architecture")

	flag.Parse()

	if *shortHelp || *longHelp {
		ansi.Blue().
			Println(usage).
			Fatal(exit.Success)
	}

	data.OpenFile("--source", sourceFile, true, os.O_RDONLY, 0644)
	data.OpenFile("--out", outFile, false, os.O_RDWR|os.O_CREATE, 0644)
	data.SetMethod(method)
	data.SetArch(arch)
}
