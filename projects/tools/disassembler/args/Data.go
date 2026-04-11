package args

import (
	"github.com/sam-caldwell/monorepo/projects/types"
	"os"
)

type Data struct {
	source *os.File
	out    *os.File
	method types.DisassemblerMethod
	arch   int
	mode   int
	debug  bool
}
