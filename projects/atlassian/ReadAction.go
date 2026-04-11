package Atlassian

import "net/http"

// ReadAction - Pattern for read operations defined in Atlassian Descriptor
type ReadAction func(*Domain, string) (*http.Request, error)
