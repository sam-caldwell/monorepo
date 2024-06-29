package matrix

import "github.com/sam-caldwell/monorepo/go/ansi"

// Print - writes the 2D Matrix to the screen.
//
// (c) 2024 Sam Caldwell.  See License.txt
func (m *Matrix) Print() {
	ansi.Blue().Printf("Matrix (%d, %d):\n")
	if m.Empty() {
		ansi.Println("<empty>").LF().LF()
	}
	for i := 0; i < m.data.Rows(); i++ {
		for j := 0; j < m.data.Cols(); j++ {
			ansi.Printf("%f ", m.data.GetDoubleAt(i, j))
		}
		ansi.LF()
	}
	ansi.LF().Reset()
}
