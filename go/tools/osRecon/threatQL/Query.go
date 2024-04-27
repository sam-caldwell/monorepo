package threatQL

import (
    "github.com/sam-caldwell/monorepo/go/fs/file"
    "github.com/sam-caldwell/monorepo/go/types"
)

// Query - A query received from the server's check-in API.
type Query struct {
	Resource types.ResourceId
	Field    types.FieldId
	Operator types.OperatorId
	Operand  types.Operand
	fileName file.File
}
