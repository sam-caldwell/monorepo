package Atlassian

import "net/http"

// CreateAction - Pattern for create operations defined in Atlassian Descriptor
type CreateAction func(*Domain) (*http.Request, error)
