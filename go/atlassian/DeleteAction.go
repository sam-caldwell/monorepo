package Atlassian

import "net/http"

// DeleteAction - Pattern for delete operations defined in Atlassian Descriptor
type DeleteAction func(*Domain, string) (*http.Request, error)
