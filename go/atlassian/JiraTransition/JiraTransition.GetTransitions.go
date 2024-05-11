package JiraTransition

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"net/http"
)

// GetTransitions - Return the transitions object by reference (raw)
func (jira *JiraTransition) GetTransitions(issueKey *string) (*Transitions, error) {

	const path = "/rest/api/3/issue/%s/transitions"

	if jira.client.GetDebug() {
		ansi.Blue().Printf("issueKey: %s", *issueKey).LF().Reset()
	}

	response, err := jira.client.Send(
		http.MethodGet,
		fmt.Sprintf(path, *issueKey),
		nil)

	if err != nil {
		return nil, err
	}

	if err = jira.Unmarshall(response); err != nil {
		return nil, err
	}

	return &jira.transitions, nil
}
