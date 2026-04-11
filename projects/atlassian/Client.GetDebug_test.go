package Atlassian

import "testing"

func TestClient_GetDebug(t *testing.T) {
	var client Client
	client.debug = true
	if !client.GetDebug() {
		t.Fatal("expect true")
	}
	client.debug = false
	if client.GetDebug() {
		t.Fatal("expect false")
	}
}
