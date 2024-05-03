package Atlassian

import "net/http"

// Descriptor - a Generic descriptor interface for working with the Atlassian API
type Descriptor interface {
    Load(fileName string) error
    Create(domain *Domain, web *http.Client) (*http.Request, error)
    Read(domain *Domain, web *http.Client) (*http.Request, error)
    Update(domain *Domain, web *http.Client) (*http.Request, error)
    Delete(domain *Domain, web *http.Client) (*http.Request, error)
    List(domain *Domain, web *http.Client) (*http.Request, error)
}
