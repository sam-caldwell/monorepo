package filters

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/__system__"
	"os"
	"runtime"
	"strings"
)

// FromCliArgs - compute our filter using our filter flags.
func (filter *Filter) FromCliArgs() (err error) {
	const notFound = -1
	if len(os.Args) < 2 {
		return nil //No flags...just go away.
	}
	hasFlag := func(flag string) (result int) {
		for pos, arg := range os.Args {
			if strings.ToLower(strings.TrimPrefix(arg, "-")) == flag {
				return pos //Found it
			}
		}
		return -1 //Not found.
	}
	hasValue := func(flag string, defaultValue string) (string, error) {
		pos := hasFlag(flag)
		if pos != notFound {
			if pos+1 < len(os.Args) {
				arg := os.Args[pos+1]
				if strings.HasPrefix(arg, "-") {
					//it's not a value.  It's a flag.
					return defaultValue, fmt.Errorf("value expected for (%s) but not found at %d. using defaults", flag, pos+1)
				}
				return arg, nil // WooHoo!  we have a value.
			}
		}
		return defaultValue, nil
	}

	filter.Commands = hasFlag("commands") != notFound
	filter.BuildEnabled = hasFlag("buildenabled") != notFound
	filter.LintEnabled = hasFlag("lintenabled") != notFound
	filter.PackEnabled = hasFlag("packenabled") != notFound
	filter.ScanEnabled = hasFlag("scanenabled") != notFound
	filter.SignEnabled = hasFlag("signenabled") != notFound
	filter.Language, err = hasValue("language", monorepo.DefaultLanguage)
	if err != nil {
		return err
	}
	filter.os, err = hasValue("os", runtime.GOOS)
	if err != nil {
		return err
	}
	filter.arch, err = hasValue("arch", runtime.GOARCH)
	return err
}
