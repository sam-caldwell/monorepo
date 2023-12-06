package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
	"time"
)

func PrintFooter(command *string, err error) {
	if err == nil {
		ansi.Green().
			LF().
			Println(strings.Repeat("=", ScreenWidth)).
			Printf("Success %s at %v\n", *command, time.Now().Format(time.RFC1123)).
			LF().
			Reset()
	} else {
		ansi.Red().
			LF().
			Println(strings.Repeat("=", ScreenWidth)).
			Printf("Failed  %s at %v", *command, time.Now().Format(time.RFC1123)).
			Printf("     Error: %v\n", err).
			LF().
			Reset()
	}
}
