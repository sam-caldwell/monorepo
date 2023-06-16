package token

import "github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"

// SetType  - Set the types.ArgTypes for the current Token
func (arg *Token) SetType(t types.ArgTypes) (err error) {

	if err := t.Valid(); err != nil {

		return err

	}

	arg.typ = t

	return nil

}
