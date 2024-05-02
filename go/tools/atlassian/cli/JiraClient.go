package cli

import (
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

// JiraClient - A struct representing a single instance of the JIRA CLI client.
type JiraClient[T Atlassian.Descriptor] struct {
    client     Atlassian.Client
    descriptor T // A JSON/YAML descriptor for an issue or project
    domain     Atlassian.Domain
    debug      bool
}
