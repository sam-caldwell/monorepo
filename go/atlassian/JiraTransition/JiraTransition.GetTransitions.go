package JiraTransition

import (
	"fmt"
	"net/http"
)

// GetTransitions - Return the transitions object by reference (raw)
func (jira *JiraTransition) GetTransitions(issueKey *string) (*Transitions, error) {

	const path = "/rest/api/3/issue/%s/transitions"

	response, err := jira.client.Send(
		http.MethodGet,
		fmt.Sprintf(path, *issueKey),
		jira.Marshall())
	if err != nil {
		return nil, err
	}
	if err = jira.Unmarshall(response); err != nil {
		return nil, err
	}
	return &jira.transitions, nil
}
