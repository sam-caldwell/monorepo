package Atlassian

import "testing"

func TestClient_GetNoop(t *testing.T) {
	var client Client
	client.noop = true
	if !client.GetNoop() {
		t.Fatal("expect true")
	}
	client.noop = false
	if client.GetNoop() {
		t.Fatal("expect false")
	}
}
