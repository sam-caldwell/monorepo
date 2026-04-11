package Atlassian

import "testing"

func TestDomain_Get(t *testing.T) {
    var d Domain
    d = Domain("testDomain")
    if v := d.Get(); v != string(d) {
        t.Fatalf("Domain.Get() mismatch")
    }

}
