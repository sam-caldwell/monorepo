package cli

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/commands"
    "net/http"
)

// Execute Given all context, execute the API action
func (client *JiraClient[T]) Execute(action commands.Commands) (err error) {

    var results *http.Request
    var web *http.Client

    if client.debug {
        ansi.Blue().Printf("action: %s", action.String()).LF().Reset()
    }

    switch action {
    case commands.Create:
        results, err = client.descriptor.Create(&client.domain, web)

    case commands.Read:
        results, err = client.descriptor.Read(&client.domain, web)

    case commands.Update:
        results, err = client.descriptor.Update(&client.domain, web)

    case commands.Delete:
        results, err = client.descriptor.Delete(&client.domain, web)

    case commands.List:
        results, err = client.descriptor.List(&client.domain, web)

    default:
        results, err = nil, fmt.Errorf("unknown/unexpected command")
    }
    if client.debug {
        ansi.Blue().Printf("err: %v", err).LF().Reset()
    }
    if err != nil {
        return err
    }

    //ToDo: better output formatting...
    if results != nil {
        ansi.Reset().Printf("%v", results).LF().Reset()
    }
    return err
}
