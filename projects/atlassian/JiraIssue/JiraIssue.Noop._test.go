package JiraIssue

import "testing"

func TestJiraIssue_Noop(t *testing.T) {
	var jira Issue
	jira.client.SetNoop(true)
	if !jira.Noop() {
		t.Fatal("noop expected true")
	}
	jira.client.SetNoop(false)
	if jira.Noop() {
		t.Fatal("noop expected false")
	}
}
