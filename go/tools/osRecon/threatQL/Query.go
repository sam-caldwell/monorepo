package threatQL

import (
	"github.com/sam-caldwell/monorepo/go/types"
)

// Query - A threat query object representing the query.
//
//	This is used in QueryFile but also this is the raw struct
//	sent across the wire between client and server.
type Query struct {
	Resource types.ResourceId `json:"Resource"`
	Field    types.FieldId    `json:"Field"`
	Operator types.OperatorId `json:"Operator"`
	Operand  types.Operand    `json:"Operand"`
}
