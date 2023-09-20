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
	"strings"
)

// GlobalFlag - Defines a global cli flag.  When the flag is detected, the provided action (function) will be called.
func (arg *ArgumentDescriptor) GlobalFlag(flags []string, description string, action func()) *ArgumentDescriptor {
	exit.OnCondition(len(flags) == 0, exit.InvalidInput, "empty flags not allowed", checkUsage)

	var helpFlag string
	for _, flag := range flags {
		if strings.HasPrefix(flag, longPrefix) {
			flagName := strings.TrimSpace(strings.ToLower(strings.TrimPrefix(flag, longPrefix)))
			exit.OnCondition(
				len(flagName) < 2,
				exit.InvalidInput,
				fmt.Sprintf("long flags must be at least two characters after '%s' (%s)", longPrefix, flagName),
				checkUsage)
			helpFlag = flag
		} else if strings.HasPrefix(flag, shortPrefix) {
			flagName := strings.TrimSpace(strings.ToLower(strings.TrimPrefix(flag, shortPrefix)))
			exit.OnCondition(
				len(flagName) != 1,
				exit.InvalidInput,
				fmt.Sprintf("short flags must be at least only one character '%s' (%s)", shortPrefix, flagName),
				checkUsage)
			helpFlag = fmt.Sprintf("%s | %s", helpFlag, flag)
		} else {
			exit.OnError(fmt.Errorf("flags must have either - or -- prefixes"), exit.InvalidInput, checkUsage)
		}
		if action == nil {
			action = func() {} //If nil, it's a noop
		}
		arg.helpText[helpFlag] = description
		// Search for and trigger the action for the global flag.
		if arg.search(flag) {
			action()
		}
	}
	return arg
}
