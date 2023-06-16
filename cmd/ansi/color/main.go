package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/ansi"
	"os"
	"strings"
)

const (
	comma = ","

	errUsage = "\n" +
		"Usage:\n" +
		"\t\tcolor <command>\n" +
		"\t\tcolor reset\n\n" +
		"\tcolors:\n" +
		"\t\t%s\n"

	errMessage        = "Error: %s%s"
	errMissingColor   = "Missing color"
	errUnknownCommand = "Unknown command"

	exitMissingColor   = 1
	exitUnknownCommand = 2

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
		commandList := strings.Join([]string{
			commandBold,
			commandDim,
			commandHidden,
			commandStrikethrough,
			commandUnderline,
			commandReset,
			commandColorBlack,
			commandColorBlue,
			commandColorCyan,
			commandColorGreen,
			commandColorMagenta,
			commandColorRed,
			commandColorWhite,
			commandColorYellow,
		}, comma)
		if condition {
			fmt.Printf(errMessage, err, fmt.Sprintf(errUsage, commandList))
			os.Exit(code)
		}
	}

	ansi.Reset()

	terminateOnError(len(os.Args) < 2, exitMissingColor, errMissingColor)

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
		terminateOnError(true, exitUnknownCommand, errUnknownCommand)
	}
}
