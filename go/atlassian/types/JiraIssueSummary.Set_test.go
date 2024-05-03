package AtlassianTypes

import "testing"

func TestJiraIssueSummary_Set(t *testing.T) {
    var summary JiraIssueSummary
    if err := summary.Set("test summary"); err != nil {
        t.Fatal(err)
    }
    if string(summary) != "test summary" {
        t.Fatal("value mismatch")
    }
}
