package cli

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/commands"
)

// Execute Given all context, execute the API action
func (client *JiraClient[T]) Execute(action commands.Commands) (err error) {

    var results any

    switch action {

    case commands.Create:
        results, err = client.descriptor.Create()

    case commands.Read:
        results, err = client.descriptor.Read()

    case commands.Update:
        results, err = client.descriptor.Update()

    case commands.Delete:
        results, err = client.descriptor.Delete()

    case commands.List:
        results, err = client.descriptor.List()

    default:
        results, err = "", fmt.Errorf("unknown/unexpected command")
    }

    if err != nil {
        return err
    }

    //ToDo: better output formatting...
    ansi.Reset().Printf("%v", results).LF().Reset()
    return err
}
