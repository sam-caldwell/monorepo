package descriptor

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/argparse/valid"
	"strings"
)

// storeLong - store a long descriptor (--string)
func (arg *Descriptor) storeLong(argument *string) (err error) {
	if strings.TrimSpace(*argument) == "" {
		arg.long = ""
		return nil
	}

	if err = valid.IsLongArg(argument); err != nil {
		return err
	}

	if *argument == argHelpLong {
		return fmt.Errorf(errReservedArg, *argument)
	}

	arg.long = strings.ToLower(strings.TrimSpace(strings.TrimSpace(*argument)))

	return err
}
