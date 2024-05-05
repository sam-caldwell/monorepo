package Atlassian

// SetDomain - Set the Atlassian tenant domain name (e.g. samcaldwell in samcaldwell.atlassian.net)
func (client *Client) SetDomain(domain *string) error {
	return client.domain.Set(domain)
}
