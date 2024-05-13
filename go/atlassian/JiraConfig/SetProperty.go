package JiraConfig

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SetProperty - Set a specific property-value on the JIRA server
func (jira *JiraConfig) SetProperty(property *string, value *string) (err error) {
	const path = "/rest/api/3/application-properties/%s"

	var response Property

	payload := []byte(fmt.Sprintf(`{
        "id": "%s",
        "value": "%s"
    }`, *property, *value))

	output, err := jira.client.Send(
		http.MethodPut,
		fmt.Sprintf(path, *property),
		payload)

	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &response)
	if err != nil {
		return err
	}
	for i, item := range jira.Properties {
		if item.Name == *property {
			jira.Properties[i] = response
			return nil
		}
	}
	//if not found, create a new property record
	jira.Properties = append(jira.Properties, response)
	return nil
}
