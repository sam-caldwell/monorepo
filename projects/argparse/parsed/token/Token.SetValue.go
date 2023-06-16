package token

// SetValue  - Set the  value for the current token
func (arg *Token) SetValue(value any) (err error) {

	//We need a typecheck before setting the value

	if err := arg.typ.Typecheck(value); err != nil {

		return err

	}

	arg.value = value

	return nil

}
