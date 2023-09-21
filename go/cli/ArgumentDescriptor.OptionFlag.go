package cli

/*
 * Copyright © 2023 Sam Caldwell <mail@samcaldwell.net>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"go/types"
	"strings"
)

func (arg *ArgumentDescriptor) OptionFlag(names []string, description string, required bool) *ArgumentDescriptor {
	var fullName string
	var helpFlag string
	if len(names) == 0 {
		panic("internal error.  empty option flag name in commandline parser definition")
	}
	for _, name := range names {
		if strings.HasPrefix(name, longPrefix) {
			fullName = strings.TrimSpace(strings.ToLower(strings.TrimPrefix(name, longPrefix)))
			exit.OnCondition(
				len(fullName) < 2,
				exit.InvalidInput,
				fmt.Sprintf("long flags must be at least two characters after '%s' (%s)", longPrefix, fullName),
				checkUsage)
			helpFlag = name
		} else if strings.HasPrefix(name, shortPrefix) {
			fullName = strings.TrimSpace(strings.ToLower(strings.TrimPrefix(name, shortPrefix)))
			exit.OnCondition(
				len(fullName) != 1,
				exit.InvalidInput,
				fmt.Sprintf("short flags must be at least only one character '%s' (%s)", shortPrefix, fullName),
				checkUsage)
			helpFlag = fmt.Sprintf("%s | %s", helpFlag, name)
		} else {
			exit.OnError(fmt.Errorf("flags must have either - or -- prefixes"), exit.InvalidInput, checkUsage)
		}
	}
	arg.helpText[helpFlag] = description

	if cmd, exists := arg.command[arg.lastCommand]; exists {
		if cmd.options == nil {
			cmd.options = make(map[string]OptionDescriptor)
		}
		cmd.options[fullName] = OptionDescriptor{
			required: required,
			vType:    types.Bool,
			value:    false,
		}
		arg.command[arg.lastCommand] = cmd
	} else {
		panic("internal error assigning action to non-existent command")
	}
	return arg
}
