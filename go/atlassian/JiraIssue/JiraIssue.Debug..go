package JiraIssue

// Debug - Return the internal client debug flag
func (jira *Issue) Debug() bool {
	return jira.client.GetDebug()
}
