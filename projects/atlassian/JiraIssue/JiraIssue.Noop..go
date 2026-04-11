package JiraIssue

// Noop - Return the internal client noop flag
func (jira *Issue) Noop() bool {
	return jira.client.GetNoop()
}
