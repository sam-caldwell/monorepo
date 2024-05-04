package AtlassianTypes

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJqlQuery_Bytes(t *testing.T) {
	// Define a sample JqlQuery instance
	jql := &JqlQuery{
		Expand:        []string{"expand1", "expand2"},
		Fields:        []string{"field1", "field2"},
		FieldsByKeys:  true,
		Jql:           "project=Atlassian",
		MaxResults:    100,
		StartAt:       0,
		ValidateQuery: "strict",
	}

	// Marshal the JqlQuery instance to JSON
	expectedBytes, err := json.Marshal(jql)
	if err != nil {
		t.Fatalf("error marshalling JqlQuery to JSON: %v", err)
	}

	// Call the Bytes method of JqlQuery
	actualBytes := jql.Bytes()

	// Compare the marshalled bytes with the result of the Bytes method
	if !reflect.DeepEqual(actualBytes, expectedBytes) {
		t.Errorf("Bytes method result does not match expected bytes.\nExpected: %s\nActual: %s", expectedBytes, actualBytes)
	}
}
