package geometry

// ColumnSize - Return the number of columns for the terminal/console (tty/pty)
func (ws *TextWindow) ColumnSize() int {
	ws.GetProperties()
	return int(ws.Cols)
}
