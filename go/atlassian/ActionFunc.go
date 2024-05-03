package Atlassian

import "net/http"

// ActionFunc - Pattern for Action functions defined in Atlassian.Descriptor
type ActionFunc func(domain *Domain) (*http.Request, error)
