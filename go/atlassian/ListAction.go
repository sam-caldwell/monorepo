package Atlassian

import (
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
)

// ListAction - Pattern for list operations defined in Atlassian Descriptor
type ListAction func(*Domain, *AtlassianTypes.JqlQuery) (*http.Request, error)
