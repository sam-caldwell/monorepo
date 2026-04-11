package AtlassianTypes

// Set - Setter for Jira issue summary
func (jira *JiraIssueSummary) Set(s string) error {
    //ToDo: Validate the issue summary
    *jira = JiraIssueSummary(s)
    return nil
}
