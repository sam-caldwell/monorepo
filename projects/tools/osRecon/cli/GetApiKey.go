package cli

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/exit"
	"github.com/sam-caldwell/monorepo/projects/simpleArgs"
	"github.com/sam-caldwell/monorepo/projects/types"
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
