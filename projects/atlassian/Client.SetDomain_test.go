package Atlassian

import "testing"

func TestClient_SetDomain(t *testing.T) {

    t.Run("happy path", func(t *testing.T) {
        var client Client
        domain := "samcaldwell"
        if err := client.SetDomain(domain); err != nil {
            t.Fatal(err)
        }
        if c := client.domain.Get(); c != domain {
            t.Fatal("expected domain mismatch")
        }
        if client.err != nil {
            t.Fatal(client.err)
        }
    })

    t.Run("sad path", func(t *testing.T) {
        var client Client
        badDomains := []string{
            "",
            "sam.caldwell",
            "sam(caldwell)",
            "sam[caldwell]",
            "sam+caldwell",
            "sam\\caldwell",
            "sam/caldwell",
            "sam*caldwell",
            "sam..caldwell",
            "sam@caldwell",
            "samcaldwell!",
            "sam!caldwell",
            "!samcaldwell",
            "sam~caldwell",
            "sam#caldwell",
            "sam$caldwell",
            "sam%caldwell",
            "sam^caldwell",
            "sam&caldwell",
            "sam<caldwell>",
            "<samcaldwell>",
            "sam?caldwell",
            "sam'caldwell",
            "sam|caldwell",
            "sam=caldwell",
        }
        for i := range badDomains {
            if err := client.SetDomain(badDomains[i]); err == nil {
                t.Fatalf("expected error on invalid domain (index: %d): '%s'", i, badDomains[i])
            } else {
                if err.Error() != MalformedJiraDomain {
                    t.Fatal("expected error message not found")
                }
            }
            if client.err == nil {
                t.Fatalf("expected internal error state on invalid domain (index: %d)", i)
            } else {
                if client.err.Error() != MalformedJiraDomain {
                    t.Fatal("expected error message not found")
                }
            }
        }
    })
}
