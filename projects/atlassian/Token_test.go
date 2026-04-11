package Atlassian

import "testing"

func TestToken_type(t *testing.T) {
    var d Token
    if d = Token("testToken"); string(d) != "testToken" {
        t.Fatal("Token failed")
    }
}
