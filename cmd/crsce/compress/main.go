package main

const (
	programName  = "compress"
	commandUsage = `
compress -h|-help
	show this screen

compress --in sourceFileName --out destinationFileName --blocksize
    compress sourceFileName and write to destinationFileName
`
)

func main() {
	//const maxBlockSize = 1048576
	//var err error
	//var c crsce.Crsce
	//var blockSize int
	//var sourceFile string
	//var destinationFile string

	//exit.OnCondition(len(os.Args) < 3, exit.MissingArg, "Missing input(s)", commandUsage)
	//exit.IfHelpRequested(commandUsage)

	//sourceFile, err = simpleArgs.GetOptionValue("--in")
	//exit.OnError(err, exit.MissingArg, commandUsage)

	//destinationFile, err = simpleArgs.GetOptionValue("--out")
	//exit.OnError(err, exit.MissingArg, commandUsage)

	//blockSize, err = simpleArgs.GetOptionIntValue("--blocksize", false)
	//exit.OnError(err, exit.MissingArg, commandUsage)
	//exit.OnCondition(blockSize < 0, exit.InvalidInput, "--blocksize must be > 0", commandUsage)
	//exit.OnCondition(blockSize > maxBlockSize, exit.InvalidInput, "--blocksize exceeds Max value", commandUsage)

	//if err = c.Open(sourceFile); err != nil {
	//	fmt.Printf("Cannot open source file.  Error: %v", err)
	//	os.Exit(exit.GeneralError)
	//}
	//if err = c.Decompress(destinationFile); err != nil {
	//	fmt.Printf("Cannot decompress to file.  Error: %v", err)
	//	os.Exit(exit.GeneralError)
	//}
}
