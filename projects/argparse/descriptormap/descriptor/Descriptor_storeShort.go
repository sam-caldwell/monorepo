package descriptor

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/argparse/valid"
	"strings"
)

// storeShort - store a short arg -{char}
func (arg *Descriptor) storeShort(argument *string) (err error) {
	if strings.TrimSpace(*argument) == "" {
		arg.short = ""
		return nil
	}

	if err = valid.IsShortArg(argument); err != nil {
		return err
	}

	if *argument == argHelpShort {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg.short = strings.ToLower(strings.TrimSpace(strings.TrimSpace(*argument)))

	return err
}
