package main

import (
	"bufio"
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
	"strings"
)

const (
	/*
	 * foreground color
	 */
	fgBlack   = "black"
	fgRed     = "red"
	fgGreen   = "green"
	fgYellow  = "yellow"
	fgBlue    = "blue"
	fgMagenta = "magenta"
	fgCyan    = "cyan"
	fgWhite   = "white"
	/*
	 * background color
	 */
	bgBlack   = "bg:black"
	bgRed     = "bg:red"
	bgGreen   = "bg:green"
	bgYellow  = "bg:yellow"
	bgBlue    = "bg:blue"
	bgMagenta = "bg:magenta"
	bgCyan    = "bg:cyan"
	bgWhite   = "bg:white"
	/*
	 * general commands
	 */
	reset = "reset"
	stdin = "stdin"
	lf    = "lf"
)

func main() {
	var message string

	fgBlackFlag := flag.Bool(fgBlack, false, "Black foreground")
	fgRedFlag := flag.Bool(fgRed, false, "Red foreground")
	fgGreenFlag := flag.Bool(fgGreen, false, "Green foreground")
	fgYellowFlag := flag.Bool(fgYellow, false, "Yellow foreground")
	fgBlueFlag := flag.Bool(fgBlue, false, "Blue foreground")
	fgMagentaFlag := flag.Bool(fgMagenta, false, "Magenta foreground")
	fgCyanFlag := flag.Bool(fgCyan, false, "Cyan foreground")
	fgWhiteFlag := flag.Bool(fgWhite, false, "White foreground")

	bgBlackFlag := flag.Bool(bgBlack, false, "Black background")
	bgRedFlag := flag.Bool(bgRed, false, "Red background")
	bgGreenFlag := flag.Bool(bgGreen, false, "Green background")
	bgYellowFlag := flag.Bool(bgYellow, false, "Yellow background")
	bgBlueFlag := flag.Bool(bgBlue, false, "Blue background")
	bgMagentaFlag := flag.Bool(bgMagenta, false, "Magenta background")
	bgCyanFlag := flag.Bool(bgCyan, false, "Cyan background")
	bgWhiteFlag := flag.Bool(bgWhite, false, "White background")

	cmdReset := flag.Bool(reset, false, "Reset the ansi color codes")
	UseStdin := flag.Bool(stdin, false, "Read from stdin")
	lineFeed := flag.Bool(lf, false, "Terminate line with line feed")

	flag.Parse()

	nonFlagArgs := flag.Args()
	if (len(nonFlagArgs) == 0) && *UseStdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message = scanner.Text()
		}
	} else {
		message = strings.TrimSpace(strings.Join(nonFlagArgs, " "))
	}

	switch true {
	case *fgBlackFlag:
		ansi.Black()
	case *fgRedFlag:
		ansi.Red()
	case *fgGreenFlag:
		ansi.Green()
	case *fgYellowFlag:
		ansi.Yellow()
	case *fgBlueFlag:
		ansi.Blue()
	case *fgMagentaFlag:
		ansi.Magenta()
	case *fgCyanFlag:
		ansi.Cyan()
	case *fgWhiteFlag:
		ansi.White()
	}

	switch true {
	case *bgBlackFlag:
		ansi.BgBlack()
	case *bgRedFlag:
		ansi.BgRed()
	case *bgGreenFlag:
		ansi.BgGreen()
	case *bgYellowFlag:
		ansi.BgYellow()
	case *bgBlueFlag:
		ansi.BgBlue()
	case *bgMagentaFlag:
		ansi.BgMagenta()
	case *bgCyanFlag:
		ansi.BgCyan()
	case *bgWhiteFlag:
		ansi.BgWhite()
	}

	ansi.Print(message)
	if *lineFeed {
		ansi.LF()
	}
	if *cmdReset {
		ansi.Reset()
	}
}
