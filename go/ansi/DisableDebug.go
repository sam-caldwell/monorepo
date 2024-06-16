package ansi

// DisableDebug - Turn off ANSI Debug print functionality. (This is a global setting)
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) DisableDebug() *Color {
	debugMode = false
	return c
}

// DisableDebug - Turn off ANSI Debug print functionality. (This is a global setting)
//
//	(c) 2023 Sam Caldwell.  MIT License
func DisableDebug() *Color {
	return (&Color{}).DisableDebug()
}
