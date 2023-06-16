ANSI Controls (Color, Navigation)
=================================

## Description

A simple color printing and cursor control library because life is not meant to be black and white.
This project allows command-line control over the cursor navigation and color without any complicated
libraries/dependencies.  It's just old-school ANSI codes doing things in Linux, MacOS, WSL the same way
we did it in DOS.  No frills, no fuss.

## Status

[![Go Tests](https://github.com/sam-caldwell/ansi/actions/workflows/go-tests.yaml/badge.svg)](https://github.com/sam-caldwell/simpleSet/actions/workflows/go-tests.yaml)

## Usage

### Installation

`go get "github.com/sam-caldwell/ansi/v2"

### Printing colors...

```
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
```

### But wait!  There's more...Navigation...

```
    ansi.Clear()  //clear the screen
```

```
    ansi.
        BgRed().Green().Blink().Bold().Print("We can move places").
        Reset().LF().LF().
        Underline().Print("Style").LF().
        Up(1).Print("Move up one").Reset().
        Left(2).Print("Move Left two").
        Hidden().Print("My secret").Reset().
        Right(3).Print("Move right three").
        Down(2).Print("Move down 2").LF().
        Strikethrough().Print("This is too much").
        Reset()

```
