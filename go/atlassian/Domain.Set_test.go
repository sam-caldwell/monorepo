package Atlassian

import "testing"

func TestJiraDomain_Set(t *testing.T) {
    var d Domain
    s := "testDomain"
    if err := d.Set(&s); err != nil {
        t.Fatal(err)
    }
    if string(d) != s {
        t.Fatal("value mismatch")
    }
}
