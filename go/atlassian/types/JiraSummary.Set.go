package AtlassianTypes

// Set - Setter for Jira issue summary
func (jira *JiraSummary) Set(s string) {
    //ToDo: Validate the issue summary
    *jira = JiraSummary(s)
}
