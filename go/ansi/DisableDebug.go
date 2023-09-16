package ansi

// DisableDebug - Turn off ANSI Debug print functionality. (This is a global setting)
func (c *Color) DisableDebug() *Color {
	debugMode = false
	return c
}

// DisableDebug - Turn off ANSI Debug print functionality. (This is a global setting)
func DisableDebug() *Color {
	c := Color{}
	c.DisableDebug()
	return &c
}
