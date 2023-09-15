package ansi

// Disable - Turn off ANSI Color codes. (This is a global setting)
func (c *Color) Disable() *Color {
	disabled = true
	return c
}

// Disable - Turn off ANSI Color codes. (This is a global setting)
func Disable() *Color {
	c := Color{}
	c.Disable()
	return &c //return the given color object
}
