package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"strings"
	"time"
)

func (m *Monorepo) PrintFooter(command *string, err error) {
	if err == nil {
		ansi.Green().
			LF().
			Println(strings.Repeat("=", ScreenWidth)).
			Printf("Success  %s at %v (elapsed: %vs)",
				*command, time.Now().Format(time.RFC1123), time.Since(m.StartTime).Seconds()).
			LF().
			Reset()
	} else {
		ansi.Red().
			LF().
			Println(strings.Repeat("=", ScreenWidth)).
			Printf("Failed  %s at %v (elapsed: %vs)",
				*command, time.Now().Format(time.RFC1123), time.Since(m.StartTime).Seconds()).
			Printf("     Error: %v\n", err).
			LF().
			Reset()
	}
}
