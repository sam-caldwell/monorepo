package Atlassian

import "net/http"

// Descriptor - a Generic descriptor interface for working with the Atlassian API
type Descriptor interface {
    Load(fileName string) error
    Create(domain *Domain) (*http.Request, error)
    Read(domain *Domain) (*http.Request, error)
    Update(domain *Domain) (*http.Request, error)
    Delete(domain *Domain) (*http.Request, error)
    List(domain *Domain) (*http.Request, error)
}

type ActionFunc func(domain *Domain) (*http.Request, error)
