package Client

// Init - Initialize the client with its basic parameters
func (jira *Client) Init(domain, apiKey string) *Client {
    jira.err = jira.domain.Set(&domain)
    if jira.err != nil {
        return jira
    }
    jira.err = jira.apiKey.FromString(&apiKey)
    return jira
}
