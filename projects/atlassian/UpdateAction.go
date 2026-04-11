package Atlassian

import "net/http"

// UpdateAction - Pattern for update operations defined in Atlassian Descriptor
type UpdateAction func(*Domain, string) (*http.Request, error)
