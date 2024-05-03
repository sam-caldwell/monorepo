package Atlassian

import "testing"

func TestClient_SetApiKey(t *testing.T) {
    var client Client

    goodToken := "ATATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-" +
        "ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=" +
        "EB457B95"

    badToken := []string{
        "ATATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-" +
            "ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=" +
            "EB457B9",
        "A(TATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-" +
            "ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=" +
            "EB45795",
        "ATATT3xFfGF0B5oiM0S_ChJv4zwN80ejCgpvs7F6nGixzASKPouaR04s7kpk8K6atrD2Jvai28Y6UDTEKnOXRT56D-" +
            "ImNrtrfLFCsUm2EoL4M4TFDxAlQqB6OKpNitXsJHYYqO1M8-mhJu8G75wo3tWkiB4hyysIH__z_GnxdXG84f-JoGAkrQ0=" +
            "EB457B950",
        "BAD_TOKEN",
    }

    if err := client.SetApiKey(goodToken); err != nil {
        t.Fatal(err)
    }

    for i := range badToken {
        bad := badToken[i]
        if err := client.SetApiKey(bad); err == nil {
            t.Fatalf("bad token %d should have failed", i)
        }
    }
}
