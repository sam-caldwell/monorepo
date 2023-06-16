package descriptor

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/counters/v2"
)

// Add - Sanitize and set the descriptor parameters.
func (arg *Descriptor) Add(pos *counters.ConditionalCounter, short string, long string,
	argType types.ArgTypes, required bool, argDefault any, help string) (err error) {

	const optionalArgumentIndicator = -1
	if err = arg.storeShort(&short); err != nil {
		return err
	}
	if err = arg.storeLong(&long); err != nil {
		return err
	}

	//Only increment position if positional argument.  otherwise store -1
	arg.pos, err = pos.IncrementIf(arg.short == "" && arg.long == "")
	if err != nil {
		panic(err)
	}

	//positional arguments cannot be flags
	if (arg.pos != optionalArgumentIndicator) && (argType == types.Flag) {
		return fmt.Errorf(errPositionalArgumentCannotBeFlag)
	}
	if err = arg.storeType(argType); err != nil {
		return err
	}
	if err = arg.storeRequired(argType, required, argDefault); err != nil {
		return err
	}
	if err = arg.storeHelpString(&help); err != nil {
		return err
	}
	return nil
}
