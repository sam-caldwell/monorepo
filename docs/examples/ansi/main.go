package main

import (
	ansi "github.com/sam-caldwell/go/v2/projects/ansi"
)

func main() {
	//ansi.Clear()
	ansi.
		Black().Print("black").LF().
		Red().Print("red|").Space().Print("|").LF().
		Green().Printf("%s", "green").LF().
		Yellow().Print("yellow|").Tab().Print("|").LF().
		Blue().Println("blue").
		Magenta().Println("magenta").
		Cyan().Println("cyan").
		White().Println("white").
		Reset().Println("reset")

	ansi.
		BgRed().Green().Blink().Bold().Print("We can move places").
		Reset().LF().LF().
		Underline().Print("Style").LF().
		Up(1).Print("Move up one").Reset().
		Left(2).Print("Move Left two").
		Hidden().Print("My secret").Reset().
		Right(3).Print("Move right three").
		Down(2).Print("Move down 2").
		Strikethrough().Print("This is too much").
		Reset()
}
