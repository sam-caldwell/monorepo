package cli

import (
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetDescriptor - Get the --descriptor value and load the file
func (client *JiraClient[T]) GetDescriptor() error {

    thisDescriptor, err := simpleArgs.GetOptionValue("--descriptor")

    if err != nil {

        return err

    }

    return client.descriptor.Load(thisDescriptor)
}
