CLI : Command-Line Interface
============================

## Description

The cli project defines reusable cobra command-line structures used by various golang projects in this monorepo.

```go
/*
 * General form...
 *      monorepo <command> <subcommand> <subcommand> <options> <action>
 */
package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/cli"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"go/types"
)

func main() {
	args := cli.NewArgParse().
		Argument("config", "cri", "check").Value("name", cli.Required, nil).Done().
		/**/
		Argument("config", "cri", "create").Value("name", cli.Required, nil).
		Option("enabled", types.Bool, cli.Optional, cli.Flag(false)).
		Option("platform", types.String, cli.Optional, cli.OptionList(words.Comma)).
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

```
