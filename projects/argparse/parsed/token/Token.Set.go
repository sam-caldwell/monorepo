package token

import "github.com/sam-caldwell/argparse/v2/argparse/types"

// Set  - Set the types.ArgTypes and value for the current Token
func (arg *Token) Set(t types.ArgTypes, v any) (err error) {

	if err := t.Typecheck(v); err != nil {

		return err

	}

	arg.typ = t

	arg.value = v

	return nil

}
