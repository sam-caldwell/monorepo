package JQL

import Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"

// JQL - Top-level struct for JQL handling
type JQL struct {
	client  *Atlassian.Client
	Queries []Query `json:"queries,omitempty"`
}
