package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
	"strings"
)

func main() {
	const (
		colorArgPosition = 1
		errUsage         = "\n" +
			"  color <command> [-n] [text:string]\n\n" +
			"  colors:\n" +
			"\t%s\n\n" +
			"  misc:\n" +
			"\t%s\n" +
			"\n-n [no new line after print, similar to bash echo]\n"
	)

	colorMap := map[string]func() *ansi.Color{
		words.Bold:          ansi.Bold,
		words.Dim:           ansi.Dim,
		words.Hidden:        ansi.Hidden,
		words.Strikethrough: ansi.Strikethrough,
		words.Underline:     ansi.Underline,
		words.Reset:         ansi.Reset,

		words.Black:   ansi.Black,
		words.Blue:    ansi.Blue,
		words.Cyan:    ansi.Cyan,
		words.Green:   ansi.Green,
		words.Magenta: ansi.Magenta,
		words.Red:     ansi.Red,
		words.White:   ansi.White,
		words.Yellow:  ansi.Yellow,
	}

	usage := func() string {
		return fmt.Sprintf(errUsage, strings.Join([]string{
			words.Black, words.Blue, words.Cyan, words.Green,
			words.Magenta, words.Red, words.White, words.Yellow,
		}, words.Comma+words.Space), strings.Join([]string{
			words.Bold, words.Dim, words.Hidden,
			words.Strikethrough, words.Underline, words.Reset,
		}, words.Comma+words.Space))
	}

	ansi.Reset()

	exit.OnCondition(len(os.Args) < 2, exit.MissingColor, exit.ErrMissingColor, usage())

	color := strings.ToLower(strings.TrimSpace(os.Args[colorArgPosition]))
	colorFunc, ok := colorMap[color]
	exit.OnCondition(!ok, exit.UnknownCommand, exit.ErrUnknownCommand, usage())
	colorFunc()

	textPos := 2
	if len(os.Args) > textPos {
		printFunc := fmt.Println
		if os.Args[textPos] == "-n" {
			printFunc = fmt.Print
			textPos++
		}
		if len(os.Args) > textPos {
			_, _ = printFunc(strings.Join(os.Args[textPos:], words.Space))
		}
	}
	defer ansi.Reset()
}
