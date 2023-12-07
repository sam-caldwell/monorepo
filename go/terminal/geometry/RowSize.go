package geometry

// RowSize - Return the number of rows for the terminal/console (tty/pty)
func (ws *TextWindow) RowSize() int {
	ws.GetProperties()
	return int(ws.Rows)
}
