package Atlassian

import "testing"

func TestClient_Init(t *testing.T) {
    var client Client

    domain := "samcaldwell"
    token := "ATATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-" +
        "ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=" +
        "EB457B95"

    if c := client.Init(domain, token); c != &client {
        t.Fatal("return address mismatch")
    }
    if client.err != nil {
        t.Fatal(client.err)
    }
    if key := client.apiKey.Get(); key != token {
        t.Fatal("token mismatch")
    }
    if d := client.domain.Get(); d != domain {
        t.Fatal("domain mismatch")
    }
}
