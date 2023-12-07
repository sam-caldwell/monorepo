package terminal

import "github.com/sam-caldwell/monorepo/go/ansi"

func BracketBox(text string, textColor string, highlight bool) *ansi.Color {
	ansi.White().Dim().Printf("[")
	switch textColor {
	case "red":
		ansi.Red()
	case "green":
		ansi.Green()
	case "yellow":
		ansi.Yellow()
	default:
		ansi.White()
	}
	if highlight {
		ansi.BgYellow().Blink()
	} else {
		ansi.BgBlack()
	}
	return ansi.Bold().Print(text).Reset().
		White().Dim().Printf("]")
}
