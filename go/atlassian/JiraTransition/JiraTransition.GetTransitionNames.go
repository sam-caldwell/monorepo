package JiraTransition

// GetTransitionNames - Given an issue key, return the map of Jira Transition names and ids for the issue.
func (jira *JiraTransition) GetTransitionNames() (result *map[string]string, err error) {
	output := make(map[string]string, 1)
	for _, item := range jira.transitions.Transitions {
		output[item.Name] = item.ID
	}
	return &output, err
}
