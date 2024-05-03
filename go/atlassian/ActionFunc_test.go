package Atlassian

import (
    "bytes"
    "fmt"
    "net/http"
    "testing"
)

func TestActionFunc(t *testing.T) {
    var d Domain
    var fn ActionFunc

    s := "test"
    if err := d.Set(&s); err != nil {
        t.Fatal(err)
    }
    fn = func(td *Domain) (*http.Request, error) {
        request, err := http.NewRequest(
            http.MethodPost,
            fmt.Sprintf(JiraUrlPattern, td.Get(), JiraApiIssue), bytes.NewBuffer([]byte("")))
        return request, err
    }
    if request, err := fn(&d); err != nil {
        t.Fatal(err)
    } else {
        if request == nil {
            t.Fatal("expected non-null request")
        }
    }
}
