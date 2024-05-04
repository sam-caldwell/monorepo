package Atlassian

import (
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
)

// Descriptor - a Generic descriptor interface for working with the Atlassian API
type Descriptor interface {
	Load(string) error
	Create(*Domain) (*http.Request, error)
	Read(*Domain, string) (*http.Request, error)
	Update(*Domain, string) (*http.Request, error)
	Delete(*Domain, string) (*http.Request, error)
	List(*Domain, *AtlassianTypes.JqlQuery) (*http.Request, error)
}
