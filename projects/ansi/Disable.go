package ansi

// Disable - Turn off ANSI Color codes. (This is a global setting)
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Disable() *Color {
	disabled = true
	return c
}

// Disable - Turn off ANSI Color codes. (This is a global setting)
//
//	(c) 2023 Sam Caldwell.  MIT License
func Disable() *Color {
	return (&Color{}).Disable()
}
