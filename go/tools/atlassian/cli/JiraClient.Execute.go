package cli

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/ansi"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "github.com/sam-caldwell/monorepo/go/misc/words"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/commands"
    "net/http"
)

// Execute Given all context, execute the API action
func (client *JiraClient[T]) Execute(action commands.Commands) (err error) {
    var (
        resp         *http.Response
        responseBody []byte
        web          *http.Client
        request      *http.Request
        actionFunc   Atlassian.ActionFunc
    )

    defer func() {
        if resp != nil {
            _ = resp.Body.Close()
        }
    }()

    if client.debug {
        ansi.Blue().Printf("action: %s", action.String()).LF().Reset()
    }

    switch action {
    case commands.Create:
        actionFunc = client.descriptor.Create
    case commands.Read:
        actionFunc = client.descriptor.Read
    case commands.Update:
        actionFunc = client.descriptor.Update
    case commands.Delete:
        actionFunc = client.descriptor.Delete
    case commands.List:
        actionFunc = client.descriptor.List
    default:
        return fmt.Errorf("unknown/unexpected command")
    }

    if request, err = actionFunc(&client.domain); err != nil {
        if client.debug {
            ansi.Blue().Printf("err: %v", err).LF().Reset()
        }
        return err
    }

    request.Header.Set(words.ContentType, words.ApplicationJson)

    if resp, err = web.Do(request); err != nil {
        return fmt.Errorf("error sending request (%v)", err)
    }

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("atlassian API responds %d (%s)", resp.StatusCode, resp.Status)
    }

    if _, err = resp.Body.Read(responseBody); err != nil {
        return fmt.Errorf("error reading response body (%v)", err)
    }

    if responseBody != nil {
        //ToDo: better output formatting...
        ansi.Reset().Printf("%v", responseBody).LF().Reset()
    }

    return err

}
