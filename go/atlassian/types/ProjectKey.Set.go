package AtlassianTypes

// Set - Setter for Jira ProjectKey
func (jira *ProjectKey) Set(s string) {
    //ToDo: Validate the issue type
    *jira = ProjectKey(s)
}
