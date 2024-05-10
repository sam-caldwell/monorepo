package JiraTransition

import (
	"fmt"
	"net/http"
)

// GetTransitionNames - Given an issue key, return the list of Jira Transition names and ids for the issue.
func (jira *JiraTransition) GetTransitionNames() (result *map[string]string, err error) {
	output := make(map[string]string, 1)
	for _, item := range jira.transitions.Transitions {
		output[item.ID] = item.Name
	}
	return &output, err
}

// GetTransitions - Return the transitions object by reference (raw)
func (jira *JiraTransition) GetTransitions(issueKey *string) *Transitions {

	const path = "/rest/api/3/issue/%s/transitions"

	response, err := jira.client.Send(
		http.MethodPost,
		fmt.Sprintf(path, *issueKey),
		jira.Marshall())
	if err != nil {
		return nil
	}
	if err = jira.Unmarshall(response); err != nil {
		return nil
	}
	return &jira.transitions
}
