package terminal

// XPixel - Return the screen size (width) in pixels
func (ws *TextWindow) XPixel() int {
	ws.GetProperties()
	return int(ws.Rows)
}
