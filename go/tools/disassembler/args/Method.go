package args

import (
	"github.com/sam-caldwell/monorepo/go/types"
)

func (data *Data) Method() types.DisassemblerMethod {
	return data.method
}
