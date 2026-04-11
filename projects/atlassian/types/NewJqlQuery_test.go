package AtlassianTypes

import (
	"testing"
)

// Mock function to simulate Jira server call for JQL validation
func mockJiraServerCall(jql string) bool {
	// Simulate Jira server validation
	return true
}

func TestNewJqlQuery(t *testing.T) {
	t.Run("Happy path test", func(t *testing.T) {
		// Happy path test
		expand := []string{"names", "schema"}
		fields := []string{"summary", "status"}
		FieldByKeys := true
		Jql := "project = 'TEST' AND assignee = currentUser()"
		MaxResults := uint(100)
		startAt := uint(0)

		jqlQuery, err := NewJqlQuery(expand, fields, FieldByKeys, Jql, MaxResults, startAt)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate the generated JqlQuery object
		if jqlQuery == nil {
			t.Error("Expected non-nil JqlQuery object, got nil")
		}

		// Add more validation if needed
	})
	t.Run("Sad path test", func(t *testing.T) {
		// Sad path test with invalid expand value
		expand := []string{"invalid_expand_value"}
		fields := []string{"summary", "status"}
		FieldByKeys := true
		Jql := "project = 'TEST' AND assignee = currentUser()"
		MaxResults := uint(100)
		startAt := uint(0)

		_, err := NewJqlQuery(expand, fields, FieldByKeys, Jql, MaxResults, startAt)

		if err == nil {
			t.Error("Expected an error for invalid expand value, got nil")
		}

		// Add more sad-path tests for other potential failure scenarios
	})
}
