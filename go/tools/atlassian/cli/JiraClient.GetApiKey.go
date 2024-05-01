package cli

import (
    "fmt"
    env "github.com/sam-caldwell/monorepo/go/environment"
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetApiKey - Get the API key from the commandline --apiKey value
func (client *JiraClient[T]) GetApiKey() (err error) {

    if argServerApiKey, err := simpleArgs.GetOptionValue("--apiKey"); err != nil {

        if err.Error() != simpleArgs.OptionNotFound {

            err = fmt.Errorf("parsing error (--apiKey): %v", err)

        } else {

            if argServerApiKey, err = env.RequireString("ATLASSIAN_TOKEN"); err != nil {

                err = fmt.Errorf("expected --apiKey or ATLASSIAN_TOKEN env var. %v", err)

            }

        }

    } else {

        err = client.client.SetApiKey(argServerApiKey)

    }

    return err

}
