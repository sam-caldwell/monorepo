package JiraConfig

import (
	"encoding/json"
	"net/http"
)

// GetFromServer - Load all property values from the server.
func (jira *JiraConfig) GetFromServer() (err error) {
	const path = "/rest/api/3/application-properties"

	var output []byte

	output, err = jira.client.Send(
		http.MethodGet,
		path,
		nil)

	if err != nil {
		return err
	}

	return json.Unmarshal(output, &jira.Properties)
}
