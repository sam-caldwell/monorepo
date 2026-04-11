package AtlassianTypes

import "testing"

func TestJiraIssueSummary_Get(t *testing.T) {
    var summary JiraIssueSummary
    summary = "test summary"
    if summary.Get() != "test summary" {
        t.Fatalf("value mismatch: '%s'", summary.Get())
    }
}
