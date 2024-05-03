package Atlassian

import "testing"

func TestJiraDomain_Get(t *testing.T) {
    var d Domain
    d = Domain("testDomain")
    if v := d.Get(); v != string(d) {
        t.Fatalf("Domain.Get() mismatch")
    }

}
