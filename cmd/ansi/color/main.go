package main

import (
	"fmt"
	ansi "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	words2 "github.com/sam-caldwell/go/v2/projects/go/misc/words"
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
			"\t-n [no new line after print, similar to bash echo]\n"
	)
	exit.IfHelpRequested(errUsage)
	exit.IfVersionRequested()

	colorMap := map[string]func() *ansi.Color{
		words2.Bold:          ansi.Bold,
		words2.Dim:           ansi.Dim,
		words2.Hidden:        ansi.Hidden,
		words2.Strikethrough: ansi.Strikethrough,
		words2.Underline:     ansi.Underline,
		words2.Reset:         ansi.Reset,

		words2.Black:   ansi.Black,
		words2.Blue:    ansi.Blue,
		words2.Cyan:    ansi.Cyan,
		words2.Green:   ansi.Green,
		words2.Magenta: ansi.Magenta,
		words2.Red:     ansi.Red,
		words2.White:   ansi.White,
		words2.Yellow:  ansi.Yellow,
	}

	usage := func() string {
		return fmt.Sprintf(errUsage, strings.Join([]string{
			words2.Black, words2.Blue, words2.Cyan, words2.Green,
			words2.Magenta, words2.Red, words2.White, words2.Yellow,
		}, words2.Comma+words2.Space), strings.Join([]string{
			words2.Bold, words2.Dim, words2.Hidden,
			words2.Strikethrough, words2.Underline, words2.Reset,
		}, words2.Comma+words2.Space))
	}

	ansi.Reset()

	exit.OnCondition(len(os.Args) < 2, exit.MissingColor, errors.MissingColor, usage())

	color := strings.ToLower(strings.TrimSpace(os.Args[colorArgPosition]))
	colorFunc, ok := colorMap[color]
	exit.OnCondition(!ok, exit.UnknownCommand, errors.UnknownCommand, usage())
	colorFunc()

	textPos := 2
	if len(os.Args) > textPos {
		printFunc := fmt.Println
		if os.Args[textPos] == "-n" {
			printFunc = fmt.Print
			textPos++
		}
		if len(os.Args) > textPos {
			_, _ = printFunc(strings.Join(os.Args[textPos:], words2.Space))
		}
	}
	defer ansi.Reset()
}
