package Atlassian

import "testing"

func TestClient_GetApiKey(t *testing.T) {
    var client Client
    client.apiKey = Token("test token")
    if client.apiKey.Get() != "test token" {
        t.Fatal("token mismatch")
    }
}
