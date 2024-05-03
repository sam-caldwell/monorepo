package Atlassian

import (
    "fmt"
    "testing"
)

func TestClientStruct(t *testing.T) {
    var client Client

    token := "foo"
    if err := client.apiKey.Set(&token); err == nil {
        t.Fatal("expected error")
    } else {
        if err.Error() != "invalid apiKey" {
            t.Fatal("mismatch error")
        }
    }
    token = "ATATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=EB457B95"
    if err := client.apiKey.Set(&token); err != nil {
        t.Fatal(err)
    } else {
        if actual := client.apiKey.Get(); actual != token {
            t.Fatal("token mismatch")
        }
    }

    domain := "bar"
    if err := client.domain.Set(&domain); err != nil {
        t.Fatal(err)
    } else {
        if domain != client.domain.Get() {
            t.Fatal("domain mismatch")
        }
    }

    testError := fmt.Errorf("test error")
    client.err = testError
    if client.err.Error() != testError.Error() {
        t.Fatal("error mismatch")
    }
}
