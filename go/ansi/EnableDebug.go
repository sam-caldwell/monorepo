package ansi

// EnableDebug - Turn on ANSI Debug print functionality. (This is a global setting)
func (c *Color) EnableDebug() *Color {
	debugMode = true
	return c
}

// EnableDebug - Turn on ANSI Debug print functionality. (This is a global setting)
func EnableDebug() *Color {
	c := Color{}
	c.EnableDebug()
	return &c
}
