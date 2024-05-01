package AtlassianTypes

import (
    "github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// Set - Setter for Jira Issue Type
func (jira *JiraIssue.JiraIssueType) Set(s string) {
    //ToDo: Validate the issue type
    *jira = JiraIssue.JiraIssueType(s)
}
