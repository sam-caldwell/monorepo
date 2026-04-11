package AtlassianTypes

// Get - Getter for Jira ProjectKey
func (jira *ProjectKey) Get() (s string) {
    return string(*jira)
}
