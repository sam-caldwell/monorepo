package ansi

// Enable - Turn on ANSI Color codes. (This is a global setting)

func (c *Color) Enable() *Color {
	disabled = false
	return c
}

// Enable - Turn on ANSI Color codes. (This is a global setting)
func Enable() *Color {
	c := Color{}
	c.Enable()
	return &c //return the given color object
}
