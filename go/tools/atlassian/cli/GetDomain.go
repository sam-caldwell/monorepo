package cli

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/exit/errors"
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetDomain Get the --domain cli parameter
func (client *JiraClient[T]) GetDomain() error {

    d, err := simpleArgs.GetOptionValue("--domain")

    if err != nil {
        return fmt.Errorf(errors.MissingArguments)
    }

    if client.debug {
        ansi.Blue().Printf("Domain: %s [%v]", d, err).LF().Reset()
    }
    return client.client.SetDomain(d)
}
