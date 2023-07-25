package simpleLogger

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"time"
)

type LogLn func(*ansi.Color, any)
type Logf func(*ansi.Color, string, ...any)
type Error func(any, int)

func Logger(quiet bool) (log LogLn,
	logf func(*ansi.Color, string, ...any), err Error) {

	if quiet {
		log = func(c *ansi.Color, message any) {}

		logf = func(c *ansi.Color, format string, args ...any) {}

		err = func(message any, exitCode int) {
			ansi.Reset().Fatal(exitCode)
		}

	} else {
		log = func(c *ansi.Color, message any) {
			c.Printf("[%v] %v\n", time.Now().UnixMilli(), message).
				Reset()
		}

		logf = func(c *ansi.Color, format string, args ...any) {
			timedArgs := append([]any{time.Now().UnixMilli()}, args...)
			c.Printf("[%v] "+format, timedArgs...).
				Reset()
		}

		err = func(message any, exitCode int) {
			ansi.Red().
				Printf("[%v] Error: '%v'\n", time.Now().UnixMilli(), message).
				Reset().
				Fatal(exitCode)
		}

	}
	return log, logf, err
}
