package token

import "github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"

// SetType  - Set the types.ArgTypes for the current Token
func (arg *Token) SetType(t types.ArgTypes) (err error) {

	if err := t.Valid(); err != nil {

		return err

	}

	arg.typ = t

	return nil

}
