package projectFilter

import "github.com/sam-caldwell/go/v2/projects/simpleArgs"

// FromCliArgs - compute our filter using our filter flags.
func (o *Filter) FromCliArgs() (err error) {
	flags := FlagList()
	for _, flag := range flags {
		//Search our os.Args for any matching flags
		var flagString string
		if flagString, err = numberToString(flag); err != nil {
			return err
		}
		if found := simpleArgs.HasFlag(flagString); found {
			//if the flag is found, perform bitwise OR to store the flag in our filter
			*o = *o | flag
		}
	}
	return nil
}
