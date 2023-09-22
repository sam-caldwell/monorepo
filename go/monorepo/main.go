package main

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
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"go/types"
)

const (
	programName        = "monorepo"
	programDescription = `A common tool for managing the monorepo and its projects.`

	DescriptionCriName = "cri name"
)

func main() {
	args := cli.NewArgParse(programName, programDescription).
		/*
		 * Global flags which may appear anywhere in the os.Args and which will be immediately processed
		 * to either update internal state or take action and terminate.
		 */
		GlobalFlag([]string{"-v", "--version"}, "show application version and exit", version.Show).
		GlobalFlag([]string{"--debug"}, "print debugging messages", cli.EnableDebug).
		GlobalFlag([]string{"--noop"}, "prevent any persistent changes", cli.EnableNoop).
		GlobalFlag([]string{"--nocolor"}, "turn off colored output", func() { ansi.Disable() }).
		/*
		 * monorepo config cri check
		 */
		Argument("validate container runtime config", words.Config, words.Cri, words.Check).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config cri create <name> [options]
		 */
		Argument("create container runtime config", words.Config, words.Cri, words.Create).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		OptionFlag([]string{"-e", "--enabled"}, "enables the object", cli.Optional).
		OptionKvList([]string{"--platform"}, "associates the object with a platform", cli.Optional, types.String, []string{}).
		Done(cli.NoAction).
		/*
		 * monorepo config cri delete <name>
		 */
		Argument("delete a container runtime config", words.Config, words.Cri, words.Delete).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config cri enable <name>
		 */
		Argument("enable a specified container runtime config", words.Config, words.Cri, words.Enable).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config cri disable <name>
		 */
		Argument("disable a specified container runtime config", words.Config, words.Cri, words.Disable).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config cri show <name>
		 */
		Argument("show a named container runtime config", words.Config, words.Cri, words.Show).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config cri list
		 */
		Argument("list all container runtime configurations", words.Config, words.Host, words.List).
		NoValue().
		Done(cli.NoAction).
		/*
		 * monorepo config host check <name>
		 */
		Argument("Verify a build-host config", words.Config, words.Host, words.Check).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config host create <name> [options]
		 */
		Argument("Create a build-host config", words.Config, words.Host, words.Create).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config host delete <name>
		 */
		Argument("Delete a build-host config", words.Config, words.Host, words.Delete).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config host enable <name>
		 */
		Argument("Enable a build-host config", words.Config, words.Host, words.Enable).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config host disable <name>
		 */
		Argument("Disable a build-host config", words.Config, words.Host, words.Disable).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config host show <name>
		 */
		Argument("Show a named build-host configuration", words.Config, words.Host, words.Show).
		Value(words.Name, DescriptionCriName, cli.Required, types.String, words.EmptyString).
		Done(cli.NoAction).
		/*
		 * monorepo config host list
		 */
		Argument("List all build-host configurations", words.Config, words.Host, words.List).
		NoValue().
		Done(cli.NoAction).
		/*
		 * monorepo -h | --help  : Show the help information and exit.  The ShowHelpIfRequested() call
		 * must be at the end before Parse() to capture all helpTest.
		 */
		ShowHelpIfRequested().
		/*
		 * Parse the actual os.Args (other than global flags already processed) and action the command.
		 */
		Parse()

	fmt.Println(args.ToString())
}
