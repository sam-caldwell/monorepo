package cli

import (
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "github.com/sam-caldwell/monorepo/go/types"
)

// JiraClient - A struct representing a single instance of the JIRA CLI client.
type JiraClient[T Atlassian.Descriptor] struct {
    client     Atlassian.Client
    apiKey     types.ApiKey // aka atlassian_token
    descriptor T            // A JSON/YAML descriptor for an issue or project
    domain     Atlassian.Domain
}
