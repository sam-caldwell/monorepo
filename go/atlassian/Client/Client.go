package Client

import (
    AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
    "github.com/sam-caldwell/monorepo/go/types"
)

type Client struct {
    apiKey types.ApiKey
    domain AtlassianTypes.AtlassianDomain
    err    error
}
