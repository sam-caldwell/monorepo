package cli

import (
    AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/Client"
    AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
    "github.com/sam-caldwell/monorepo/go/types"
)

// JiraClient - A struct representing a single instance of the JIRA CLI client.
type JiraClient[T AtlassianTypes.Descriptor] struct {
    client     AtlassianJira.Client
    apiKey     types.ApiKey // aka atlassian_token
    descriptor T            // A JSON/YAML descriptor for an issue or project
    domain     AtlassianTypes.AtlassianDomain
}
