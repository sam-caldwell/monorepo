package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
	"strings"
)

const (
	errUsage = "\t\tcolor <command>\n" +
		"\t\tcolor reset\n\n" +
		"\tcolors:\n" +
		"\t\t%s\n\n" +
		"\tmisc:\n" +
		"\t\t%s\n"

	commandBold          = "bold"
	commandDim           = "dim"
	commandHidden        = "hidden"
	commandStrikethrough = "strikethrough"
	commandUnderline     = "underline"

	commandReset = "reset"

	commandColorBlack   = "black"
	commandColorBlue    = "blue"
	commandColorCyan    = "cyan"
	commandColorGreen   = "green"
	commandColorMagenta = "magenta"
	commandColorRed     = "red"
	commandColorWhite   = "white"
	commandColorYellow  = "yellow"
)

func main() {
	terminateOnError := func(condition bool, code int, err string) {
		exit.OnCondition(
			condition,
			code, err,
			fmt.Sprintf(errUsage, strings.Join([]string{
				commandColorBlack,
				commandColorBlue,
				commandColorCyan,
				commandColorGreen,
				commandColorMagenta,
				commandColorRed,
				commandColorWhite,
				commandColorYellow,
			}, words.Comma), strings.Join([]string{
				commandBold,
				commandDim,
				commandHidden,
				commandStrikethrough,
				commandUnderline,
			}, words.Comma)))
	}

	ansi.Reset()

	terminateOnError(len(os.Args) < 2, exit.MissingColor, exit.ErrMissingColor)

	switch strings.ToLower(strings.TrimSpace(os.Args[1])) {
	case commandBold:
		ansi.Bold()
	case commandDim:
		ansi.Dim()
	case commandHidden:
		ansi.Hidden()
	case commandStrikethrough:
		ansi.Strikethrough()
	case commandUnderline:
		ansi.Underline()
	case commandReset:
		ansi.Reset()

	case commandColorBlack:
		ansi.Black()
	case commandColorBlue:
		ansi.Blue()
	case commandColorCyan:
		ansi.Cyan()
	case commandColorGreen:
		ansi.Green()
	case commandColorMagenta:
		ansi.Magenta()
	case commandColorRed:
		ansi.Red()
	case commandColorWhite:
		ansi.White()
	case commandColorYellow:
		ansi.Yellow()

	default:
		terminateOnError(true, exit.UnknownCommand, exit.ErrUnknownCommand)
	}
}
