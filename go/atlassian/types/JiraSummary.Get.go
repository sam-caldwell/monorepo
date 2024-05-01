package AtlassianTypes

// Get - Getter for Jira issue summary
func (jira *JiraSummary) Get() (s string) {
    return string(*jira)
}
