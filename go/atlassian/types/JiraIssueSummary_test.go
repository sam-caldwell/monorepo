package AtlassianTypes

import "testing"

func TestJiraIssueSummary(t *testing.T) {
    var summary JiraIssueSummary
    summary = JiraIssueSummary("test summary")
    if string(summary) != "test summary" {
        t.Fatal("value mismatch")
    }
}
