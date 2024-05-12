package JiraIssue

import (
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
)

// Issue - a structure representing a JIRA issue
type Issue struct {
	IssueKey string `json:"issueKey,omitempty"`
	Fields   struct {
		Project struct {
			Key string `json:"key,omitempty"`
		} `json:"project,omitempty"`
		Summary     string `json:"summary,omitempty"`
		Description struct {
			Type    string `json:"type,omitempty"`
			Version int    `json:"version,omitempty"`
			Content []struct {
				Type    string `json:"type,omitempty"`
				Content []struct {
					Type string `json:"type,omitempty"`
					Text string `json:"text,omitempty"`
				} `json:"content,omitempty"`
			} `json:"content,omitempty"`
		} `json:"description,omitempty"`
		IssueType struct {
			Name string `json:"name,omitempty"`
		} `json:"issuetype,omitempty"`
	} `json:"fields,omitempty"`
	client Atlassian.Client
	jql    AtlassianTypes.JqlQuery
}
