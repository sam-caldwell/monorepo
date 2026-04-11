package main

import (
	"github.com/sam-caldwell/monorepo/projects/exit"
	"github.com/sam-caldwell/monorepo/projects/simpleArgs"
	"github.com/sam-caldwell/monorepo/projects/tools/atlassian/JiraActions"
)

func main() {

	command := simpleArgs.GetCommand("command (create,read,update,delete,list,transition")

	object := simpleArgs.GetCommand("object (issue,project)")

	exit.TerminateOnError(JiraActions.Router(&command, &object))

}
