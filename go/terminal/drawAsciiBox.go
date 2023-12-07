package terminal

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// DrawAsciiBox - Draw an ASCII text box of a given width
func DrawAsciiBox(color *ansi.Color, text []string, width int) {
	//screenWidth := ColumnSize()
	//adjustment := 4
	barWidth := width - 2
	color.LF().
		Print("╔").Printf(strings.Repeat("═", barWidth)).Print("╗").LF()
	for _, line := range text {
		lineWidth := len(line)
		paddingWidth := barWidth - (lineWidth + 2)
		padding := strings.Repeat(words.Space, paddingWidth)
		color.Printf("║ %s%s ║", line, padding).LF()
	}
	color.Printf("╚%s╝", strings.Repeat("═", barWidth)).LF().Reset()
}
