package cli

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetDescriptor - Get the --descriptor value and load the file
func (client *JiraClient[T]) GetDescriptor() error {

    thisDescriptor, err := simpleArgs.GetOptionValue("--descriptor")

    if err != nil {

        return err

    }

    if client.debug {
        ansi.Blue().Printf("GetDescriptor(): %s", thisDescriptor).LF().Reset()
    }
    return client.descriptor.Load(thisDescriptor)
}
