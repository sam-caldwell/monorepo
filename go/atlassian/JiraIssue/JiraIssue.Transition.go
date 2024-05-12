package JiraIssue

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraTransition"
	"net/http"
)

// Transition - Transition workflow step for a given ticket
func (jira *Issue) Transition(step *string) (output []byte, err error) {

	const path = "/rest/api/3/issue/%s/transitions"

	var transition JiraTransition.JiraTransition
	var steps *map[string]string

	if err = transition.Init(&jira.client, &jira.IssueKey); err != nil {
		return nil, err
	}

	if steps, err = transition.GetTransitionNames(); err != nil {
		return nil, err
	}

	if stepId, ok := (*steps)[*step]; ok {

		//var action JiraTransition.TransitionAction
		//action.Transition.ID = stepId

		payload := []byte(fmt.Sprintf("{\"transition\":\"%s\"}", stepId))

		if jira.client.GetDebug() {
			ansi.Blue().
				Printf("step target %s [%s]", stepId, *step).LF().
				Printf("payload: %s", payload).LF().Reset()
		}

		output, err = jira.client.Send(
			http.MethodPost,
			fmt.Sprintf(path, jira.IssueKey),
			payload)

	} else {
		if jira.client.GetDebug() {
			ansi.Blue().Printf("Searching for %s", *step).LF().Reset()
			for k, v := range *steps {
				ansi.Blue().Printf("{\"id\":\"%s\", \"name\":\"%s\"}", k, v).LF().Reset()
			}
		}

		return nil, fmt.Errorf("transition step %s not found in workflow", *step)

	}
	return output, err
}
