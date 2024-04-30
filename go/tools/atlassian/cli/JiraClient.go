package cli

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types"
)

// JiraClient - A struct representing a single instance of the JIRA CLI client.
type JiraClient struct {
	apiKey     types.ApiKey
	command    Commands
	object     Objects
	descriptor file.File
}
