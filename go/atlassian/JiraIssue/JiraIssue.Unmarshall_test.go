package JiraIssue

import (
	"reflect"
	"testing"
)

func TestIssueUnmarshall(t *testing.T) {
	// Define a sample JSON data
	jsonData := []byte(`{
		"issueKey": "ISSUE-123",
		"fields": {
			"project": {
				"key": "PROJECTKEY"
			},
			"summary": "Sample summary",
			"description": {
				"type": "description",
				"version": 1,
				"content": [
					{
						"type": "paragraph",
						"content": [
							{
								"type": "text",
								"text": "Sample description"
							}
						]
					}
				]
			},
			"issuetype": {
				"name": "Task"
			}
		}
	}`)

	// Define expected structure after unmarshalling
	expected := Issue{
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

	// Create an empty issue to unmarshal into
	var issue Issue

	// Unmarshal the JSON data into the issue
	err := issue.Unmarshall(jsonData)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Compare expected and actual structures
	if !reflect.DeepEqual(expected, issue) {
		t.Fatalf("Expected:\n%+v\n\nActual:\n%+v", expected, issue)
	}
}
