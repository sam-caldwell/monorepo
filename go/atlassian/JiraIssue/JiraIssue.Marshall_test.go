package JiraIssue

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestIssueMarshall(t *testing.T) {
	// Define a sample issue
	issue := Issue{
		IssueKey: "ISSUE-123",
		Fields: struct {
			Project struct {
				Key string `json:"key"`
			} `json:"project"`
			Summary     string `json:"summary"`
			Description struct {
				Type    string `json:"type"`
				Version int    `json:"version"`
				Content []struct {
					Type    string `json:"type"`
					Content []struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"content"`
				} `json:"content"`
			} `json:"description"`
			IssueType struct {
				Name string `json:"name"`
			} `json:"issuetype"`
		}{
			Project: struct {
				Key string `json:"key"`
			}{
				Key: "PROJECTKEY",
			},
			Summary: "Sample summary",
			Description: struct {
				Type    string `json:"type"`
				Version int    `json:"version"`
				Content []struct {
					Type    string `json:"type"`
					Content []struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"content"`
				} `json:"content"`
			}{
				Type:    "description",
				Version: 1,
				Content: []struct {
					Type    string `json:"type"`
					Content []struct {
						Type string `json:"type"`
						Text string `json:"text"`
					} `json:"content"`
				}{
					{
						Type: "paragraph",
						Content: []struct {
							Type string `json:"type"`
							Text string `json:"text"`
						}{
							{
								Type: "text",
								Text: "Sample description",
							},
						},
					},
				},
			},
			IssueType: struct {
				Name string `json:"name"`
			}{
				Name: "Task",
			},
		},
	}

	// Marshal the issue
	marshaledIssue := issue.Marshall()

	// Unmarshal the marshaled JSON data
	var unmarshaledIssue Issue
	err := json.Unmarshal(marshaledIssue, &unmarshaledIssue)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Compare the original and unmarshaled issues
	if !reflect.DeepEqual(issue, unmarshaledIssue) {
		t.Fatalf("Expected:\n%+v\n\nActual:\n%+v", issue, unmarshaledIssue)
	}
}
