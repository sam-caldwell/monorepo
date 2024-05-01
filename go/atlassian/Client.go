package Atlassian

import (
    "github.com/sam-caldwell/monorepo/go/types"
)

type Client struct {
    apiKey types.ApiKey
    domain Domain
    err    error
}
