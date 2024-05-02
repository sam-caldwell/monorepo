package cli

import (
    "fmt"
    "github.com/sam-caldwell/monorepo/go/exit/errors"
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetDomain Get the --domain cli parameter
func (client *JiraClient[T]) GetDomain() error {

    thisDescriptor, err := simpleArgs.GetOptionValue("--domain")

    if err != nil {
        return fmt.Errorf(errors.MissingArguments)
    }

    return client.client.SetDomain(thisDescriptor)
}
