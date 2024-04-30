package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	env "github.com/sam-caldwell/monorepo/go/environment"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
	"github.com/sam-caldwell/monorepo/go/types"
)

// GetApiKey - Get the API key from the commandline --apiKey value
func GetApiKey() types.ApiKey {

	argServerApiKey, err := simpleArgs.GetOptionValue("--apiKey")
	if err != nil {
		if err.Error() != simpleArgs.OptionNotFound {
			ansi.Red().
				Printf("Parsing error (%s): %v", "--apiKey", err.Error()).
				LF().
				Fatal(exit.GeneralError).
				Reset()
		} else {
			if argServerApiKey, err = env.RequireString("ATLASSIAN_API_KEY"); err != nil {
				ansi.Red().
					Printf("Expected --apiKey or ATLASSIAN_API_KEY env var. %v", err).
					LF().
					Fatal(exit.GeneralError).
					Reset()
			}
		}
	}

	var serverApiKey types.ApiKey

	exit.TerminateOnError(serverApiKey.FromString(&argServerApiKey))

	return serverApiKey

}
