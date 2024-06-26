package Atlassian

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"net/http"
	"testing"
)

func TestClient_SetAuthHeader(t *testing.T) {
	var client Client

	domain := "samcaldwell"
	token := "ATATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-" +
		"ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=" +
		"EB457B95"

	_ = client.Init(domain, token)

	if client.err != nil {
		t.Fatal(client.err)
	}

	request, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(JiraUrlPattern, client.domain.Get(), atlassianFqdn, JiraApiIssue), bytes.NewBuffer([]byte("")))
	if err != nil {
		t.Fatal(err)
	}

	if err := client.SetAuthHeader(request); err != nil {
		t.Fatal(err)
	}

	if a := request.Header.Get(words.Authorization); a != words.Basic+client.GetApiKey() {
		t.Fatal("authorization header mismatch")
	}

}
