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
package main

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"go/types"
)

const (
	programName        = "monorepo"
	programDescription = `A common tool for managing the monorepo and its projects.`
)

func noopAction() func() {
	// This is a placeholder for actions
	return func() {}
}

func main() {
	argument := cli.NewArgumentParser(programName, programDescription).
		VersionFlag().
		NoopFlag().
		DebugFlag().
		SubCommand("build", noopAction()).
		SubCommandWithChildren("config", cli.
			SubCommandWithChildren("cri", cli.
				SubCommandWithOptions("create", cli.
					Value("name", types.String).
						OptionFlag("enabled", types.Bool, false, cli.Required).
						OptionKeyValue("description", types.String, words.EmptyString, cli.Required).
						OptionKeyValue("platform", types.String, words.EmptyString, cli.Required)))).
				EndSubCommandWithOptions().
			SubCommandWithChildren("host")
		SubCommand("exec", noopAction()).
		SubCommand("lint", noopAction()).
		SubCommand("scan", noopAction()).
		SubCommand("sign", noopAction()).
		SubCommand("test", noopAction()).
}
