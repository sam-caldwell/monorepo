package JiraIssue

// Get - Getter for Jira Issue Type
func (jira *JiraIssueType) Get() (s string) {
    return string(*jira)
}
