package Atlassian

import (
	"bytes"
	"fmt"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
	"testing"
)

func TestListAction(t *testing.T) {
	var d Domain
	var fn ListAction

	s := "test"
	if err := d.Set(&s); err != nil {
		t.Fatal(err)
	}
	fn = func(td *Domain, jql *AtlassianTypes.JqlQuery) (*http.Request, error) {
		request, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf(JiraUrlPattern, td.Get(), atlassianFqdn, JiraApiIssue), bytes.NewBuffer([]byte("")))
		return request, err
	}
	if request, err := fn(&d, nil); err != nil {
		t.Fatal(err)
	} else {
		if request == nil {
			t.Fatal("expected non-null request")
		}
	}
}
