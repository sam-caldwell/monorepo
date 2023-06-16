package argparse

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/descriptormap/descriptor"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/argparse/v2/argparse/valid"
	"log"
	"os"
	"strings"
)

// Parse - Parse os.Args
func (arg *Arguments) Parse() *Arguments {

	var optionalArgEncountered = false
	var valueExpected = false
	var thisName *string
	var thisDescriptor *descriptor.Descriptor = nil

	if err := arg.descriptors.FindDuplicates(); err != nil {
		_ = arg.err.Push(err)
		return arg
	}
	for position, rawToken := range os.Args[1:] {
		log.Printf("start parser for (position:%d, token: %s):", position, rawToken)

		// clean the token any intentional whitespace should be wrapped in quotes.
		token := strings.TrimSpace(rawToken)

		// we expect a value and will not consider -a or --all patterns
		if valueExpected {
			log.Printf("capture value(%v) for argument (%s)", token, *thisName)
			var value any = token //ToDo: Cast the value to the proper type.
			typ := thisDescriptor.GetType()
			value = token
			if err := arg.results.Add(thisName, typ, &value); err != nil {
				log.Printf("failed to add argument for position %d (%s) '%v'", position, *thisName, value)
				_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
			}
			thisName = nil
			thisDescriptor = nil
			valueExpected = false
			log.Printf("---value is stored---")
			continue
		}
		// Deal with the long arguments...
		if err := valid.IsLongArg(&token); err == nil {
			log.Printf("Found long arg %s", token)
			thisName, thisDescriptor = arg.descriptors.GetByLong(&token)
			if thisName == nil {
				_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
				return arg //bail!
			}
			//Flags have no values.  We can just store them.
			if typ := thisDescriptor.GetType(); typ == types.Flag {
				var value any = true
				_ = arg.err.Push(arg.results.Add(thisName, types.Flag, &value))
				valueExpected = false
			} else {
				valueExpected = true //we need the next item to be a value
			}
			optionalArgEncountered = true
			continue //Just go get the next one
		}

		//Deal with the short arguments
		if err := valid.IsShortArg(&token); err == nil {
			log.Printf("Found short arg %s", token)
			thisName, thisDescriptor = arg.descriptors.GetByShort(&token)
			if thisName == nil {
				_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
				return arg //bail!
			}
			//Flags have no values.  We can just store them.
			if typ := thisDescriptor.GetType(); typ == types.Flag {
				var value any = true
				_ = arg.err.Push(arg.results.Add(thisName, types.Flag, &value))
				valueExpected = false
			} else {
				valueExpected = true //we need the next item to be a value
			}
			optionalArgEncountered = true
			continue //Just go get the next one
		}

		//Deal with the positional arguments (token is a value now)
		if optionalArgEncountered {
			//If we have already encountered a short -a or long --all optional argument, bail.
			_ = arg.err.Push(fmt.Errorf(errPositionalArgumentFollowingOptional, position))
			return arg
		}

		//Otherwise look up our expected argument by position.
		thisName, thisDescriptor = arg.descriptors.GetByPosition(position)
		log.Printf("GetByPosition(%d) %v %d", position, *thisName, thisDescriptor.GetPosition())
		//But if the result (thisName) is nil, we weren't expecting an arg at this position.  bail.
		if thisName == nil {
			log.Printf("No argument found at position %d", position)
			// if position not found in our descriptor map then we have a syntax error.
			_ = arg.err.Push(fmt.Errorf(errFoundUndefinedArgument, token))
			return arg
		}

		//However, if all this works out, token must be a positional value.  Type-check it. Bail on fail.
		var value any = token //ToDo: Cast the value to the proper type.
		typ := thisDescriptor.GetType()
		if err := typ.Typecheck(token); err != nil {
			log.Printf("typecheck failed on token %s (value: %v", *thisName, value)
			_ = arg.err.Push(err)
			return arg
		}

		// Store this value
		if err := arg.results.Add(thisName, typ, &value); err != nil {
			log.Printf("Failed to store positional token %s (value: %v", *thisName, value)
			_ = arg.err.Push(err)
			return arg
		}

		// Set our state to expect a new token
		valueExpected = false

		continue
	} /*end of for*/
	//ToDo: ensure all required arguments are satisfied.
	return arg
}
