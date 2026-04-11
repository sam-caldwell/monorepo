package Atlassian

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestReadAction(t *testing.T) {
	var d Domain
	var fn ReadAction

	s := "test"
	if err := d.Set(&s); err != nil {
		t.Fatal(err)
	}
	fn = func(td *Domain, n string) (*http.Request, error) {
		request, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf(JiraUrlPattern, td.Get(), atlassianFqdn, JiraApiIssue), bytes.NewBuffer([]byte("")))
		return request, err
	}
	if request, err := fn(&d, ""); err != nil {
		t.Fatal(err)
	} else {
		if request == nil {
			t.Fatal("expected non-null request")
		}
	}
}
