package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
	"github.com/sam-caldwell/monorepo/go/types"
)

// GetApiKey - Get the API key from the commandline --apiKey value
func GetApiKey() types.ApiKey {

	argServerApiKey, err := simpleArgs.GetOptionValue("--apiKey")
	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", "--apiKey", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	var serverApiKey types.ApiKey

	exit.TerminateOnError(serverApiKey.FromString(&argServerApiKey))

	return serverApiKey

}
