package terminal

// YPixel - Return the screen size (height) in pixels
func (ws *TextWindow) YPixel() int {
	ws.GetProperties()
	return int(ws.Rows)
}
