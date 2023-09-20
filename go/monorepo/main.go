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
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"go/types"
)

const (
	programName        = "monorepo"
	programDescription = `A common tool for managing the monorepo and its projects.`
)

func main() {
	args := cli.NewArgParse(programName, programDescription).
		GlobalFlag([]string{"-v", "--version"}, "show application version", version.Show).
		GlobalFlag([]string{"--debug"}, "print debugging messages", cli.EnableDebug).
		GlobalFlag([]string{"--noop"}, "prevent any persistent changes", cli.EnableNoop).
		/*
		 * monorepo config cri check
		 */
		Argument("config", "cri", "check").Value("name", cli.Required, nil).Done().
		/*
		 * monorepo config cri create <name> [options]
		 */
		Argument("config", "cri", "create").
		Value("name", cli.Required, nil).
		Option([]string{"--enabled"}, "enables the object", types.Bool, cli.Optional, cli.Flag(false)).
		Option([]string{"--platform"}, "associates the object with a platform", types.String, cli.Optional, cli.OptionList(words.Comma)).
		Done().
		/**/
		Argument("config", "cri", "delete").Value("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "enable").Value("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "disable").Value("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "show").Value("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "list").Value(nil, cli.None, nil).Done().
		/**/
		Argument("config", "cri", "check").StringValue("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "create").StringValue("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "delete").v("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "enable").StringValue("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "disable").StringValue("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "show").StringValue("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "list").NoValue().Done().
		/**/
		Parse()

	fmt.Println(args.ToString())
}
