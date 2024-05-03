package Atlassian

import (
	"fmt"
	"testing"
)

func TestJiraUrlFactory(t *testing.T) {
	tenant := "samcaldwell"
	apiPath := "/a/path/we/should/have"

	expected := fmt.Sprintf(JiraUrlPattern, tenant, atlassianFqdn, apiPath)
	// JiraUrlFactory - form the JIRA URL
	actual := JiraUrlFactory(JiraUrlPattern, tenant, apiPath)

	if actual != expected {
		t.Fatalf("url mismatch.\n"+
			"\tactual:%s\n"+
			"\texpected:%s\n", actual, expected)
	}
}
