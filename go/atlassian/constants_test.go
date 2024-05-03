package Atlassian

import "testing"

func TestConstants_DefaultJqLQuery(t *testing.T) {
	if DefaultJqlQuery != "*" {
		t.Fatal("DefaultJqlQuery mismatch")
	}
	if atlassianTokenRegex != `^[A-Za-z0-9+/\-=_]{192}$` {
		t.Fatal("atlassianTokenRegex mismatch")
	}
	if MalformedJiraDomain != "JiraDomain malformed" {
		t.Fatal("JiraDomain mismatch")
	}
	if JiraDomainRegEx != `^[a-zA-Z][a-zA-Z0-9\-_]+[a-zA-Z0-9]$` {
		t.Fatal("JiraDomainRegEx mismatch")
	}
	if JiraDomainNotInitialized != "JiraDomain not initialized" {
		t.Fatal("JiraDomainNotInitialized mismatch")
	}
	if JiraUrlPattern != "https://%s.%s/%s" {
		t.Fatal("JiraUrlPattern mismatch")
	}
	if JiraApiIssue != "/rest/api/2/issue/" {
		t.Fatal("JiraApiIssue mismatch")
	}
}
