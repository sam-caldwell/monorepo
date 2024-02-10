package args

import (
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
)

type Data struct {
	source *os.File
	out    *os.File
	method types.DisassemblerMethod
	arch   int
	mode   int
}
