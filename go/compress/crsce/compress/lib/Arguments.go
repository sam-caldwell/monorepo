package crsce

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
