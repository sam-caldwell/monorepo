package terminal

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

func ensureLinesFit(width int, text []string) []string {

	var result []string

	if width <= 0 {
		width = 20
	}
	for _, line := range text {
		line = strings.Replace(line, "\n", " ", -1)
		result = append(result, line)
	}
	text = result
	result = []string{}
	for _, line := range text {
		line = strings.TrimSpace(line)
		if len(line) < width {
			result = append(result, line)
		} else {
			lines := SplitStringIntoLines(line, width)
			for _, ln := range lines {
				ln = strings.TrimSpace(ln)
				result = append(result, ln)
			}
		}
	}
	return result
}

// DrawAsciiBox - Draw an ASCII text box of a given width
func DrawAsciiBox(color *ansi.Color, text []string, width int) {
	text = ensureLinesFit(width-10, text)

	barWidth := width - 2

	color.LF().Print("╔").Printf(strings.Repeat("═", barWidth)).Print("╗").LF()

	for _, line := range text {
		//line = strings.TrimSpace(line)

		lineWidth := len(line)

		paddingWidth := barWidth - (lineWidth + 2)

		padding := strings.Repeat(words.Space, paddingWidth)

		color.Printf("║ %s%s ║", line, padding).LF()

	}

	color.Printf("╚%s╝", strings.Repeat("═", barWidth)).LF().Reset()
}
