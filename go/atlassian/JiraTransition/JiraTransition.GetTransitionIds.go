package JiraTransition

// GetTransitionIds - Given an issue key, return the map of Jira Transition ids an related data.
func (jira *JiraTransition) GetTransitionIds() (result *map[string]string, err error) {
	output := make(map[string]string, 1)
	for _, item := range jira.transitions.Transitions {
		output[item.ID] = item.Name
	}
	return &output, err
}
