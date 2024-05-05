package Atlassian

// Client - General client for Atlassian API calls
type Client struct {
	apiKey Token
	domain Domain
	debug  bool
	noop   bool
}
