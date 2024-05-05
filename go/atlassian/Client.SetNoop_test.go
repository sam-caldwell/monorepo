package Atlassian

import "testing"

func TestClient_SetNoop(t *testing.T) {

	var client Client

	client.SetNoop(true)

	if !client.noop {
		t.Fatal("expect true")
	}

	client.SetNoop(false)

	if client.noop {
		t.Fatal("expect false")
	}

}
