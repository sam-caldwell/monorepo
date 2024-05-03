package Atlassian

import "testing"

func TestDomainType(t *testing.T) {
    var d Domain
    if d = Domain("testDomain"); string(d) != "testDomain" {
        t.Fatal("domain failed")
    }
}
