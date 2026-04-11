package JiraIssue

import "testing"

func TestJiraIssue_Debug(t *testing.T) {
	var jira Issue
	jira.client.SetDebug(true)
	if !jira.Debug() {
		t.Fatal("debug expected true")
	}
	jira.client.SetDebug(false)
	if jira.Debug() {
		t.Fatal("debug expected false")
	}
}
