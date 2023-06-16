package vartype

// GetAny - blindly return the data from the object.  (error is always nil, but it exists for future use)
func (o *Generic) GetAny() (data any, err error) {
	return o.data, nil
}
