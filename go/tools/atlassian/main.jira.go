package main

import (
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
	"github.com/sam-caldwell/monorepo/go/tools/atlassian/JiraActions"
)

func main() {

	command := simpleArgs.GetCommand("command (create,read,update,delete,list")
	object := simpleArgs.GetCommand("object (issue,project)")
	exit.TerminateOnError(JiraActions.Router(&command, &object))

}
