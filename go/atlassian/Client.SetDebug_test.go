package Atlassian

import "testing"

func TestClient_SetDebug(t *testing.T) {

	var client Client

	client.SetDebug(true)

	if !client.debug {
		t.Fatal("expect true")
	}

	client.SetDebug(false)

	if client.debug {
		t.Fatal("expect false")
	}

}
