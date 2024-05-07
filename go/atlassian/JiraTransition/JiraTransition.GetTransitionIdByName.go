package JiraTransition

import (
	"fmt"
	"net/http"
)

// GetTransitionIdByName - Given a workflow step name and assoc. issue key, return the step name
func (jira *JiraTransition) GetTransitionIdByName(IssueKey, stepName *string) error {

	const path = "/rest/api/3/issue/%s/transitions"

	response, err := jira.client.Send(
		http.MethodPost,
		fmt.Sprintf(path, *IssueKey),
		transition.JsonMarshall())
	if err != nil {
		return err
	}
	if err = jira.Unmarshall(response); err != nil {
		return err
	}
}
