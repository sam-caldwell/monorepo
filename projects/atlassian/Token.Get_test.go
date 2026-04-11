package Atlassian

import "testing"

func TestToken_Get(t *testing.T) {
    var d Token

    d = Token("testToken")

    if v := d.Get(); v != string(d) {
        t.Fatalf("Token.Get() testToken")
    }

}
