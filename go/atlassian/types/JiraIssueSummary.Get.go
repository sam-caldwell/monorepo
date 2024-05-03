package AtlassianTypes

// Get - Getter for Jira issue summary
func (jira *JiraIssueSummary) Get() (s string) {
    return string(*jira)
}
