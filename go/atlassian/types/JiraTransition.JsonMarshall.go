package AtlassianTypes

import "encoding/json"

// Marshall - marshall the jira transition request and return its json []byte form
func (jira *JiraTransition) Marshall() []byte {

	result, err := json.Marshal(jira)

	if err != nil {
		panic("internal error")
	}

	return result

}
