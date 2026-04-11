package ansi

// Enable - Turn on ANSI Color codes. (This is a global setting)
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Enable() *Color {
	disabled = false
	return c
}

// Enable - Turn on ANSI Color codes. (This is a global setting)
//
//	(c) 2023 Sam Caldwell.  MIT License
func Enable() *Color {
	return (&Color{}).Enable()
}
