package JiraTransition

// GetTransitionIds - Given an issue key, return the map of Jira Transition ids an related data.
func (jira *JiraTransition) GetTransitionIds() (result *map[string]Transition, err error) {
	output := make(map[string]Transition, 1)
	for _, item := range jira.transitions.Transitions {
		output[item.Name] = item
	}
	return &output, err
}
