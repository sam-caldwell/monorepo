package cli

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/ansi"
    env "github.com/sam-caldwell/monorepo/go/environment"
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetApiKey - Get the API key from the commandline --apiKey value
func (client *JiraClient[T]) GetApiKey() (err error) {

    argServerApiKey, err := simpleArgs.GetOptionValue("--apiKey")

    if err != nil {
        if err.Error() == simpleArgs.OptionNotFound {
            if argServerApiKey, err = env.RequireString("ATLASSIAN_TOKEN"); err != nil {
                err = fmt.Errorf("expected --apiKey or ATLASSIAN_TOKEN env var. %v", err)
            }
        }
    }

    err = client.client.SetApiKey(argServerApiKey)

    if client.debug {
        ansi.Blue().
            Printf("ApiKey: '%s' "+
                "error: %v", client.client.GetApiKey(), err).LF().
            Reset()
    }

    return err

}
