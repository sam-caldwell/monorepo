package args

import (
	"github.com/sam-caldwell/monorepo/projects/types"
)

func (data *Data) Method() types.DisassemblerMethod {
	return data.method
}
