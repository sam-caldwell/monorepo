package JiraIssue

// Set - Setter for Jira Issue Type
func (jira *JiraIssueType) Set(s string) {
    //ToDo: Validate the issue type
    *jira = JiraIssueType(s)
}
