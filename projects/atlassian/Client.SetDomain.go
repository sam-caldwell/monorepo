package Atlassian

import "fmt"

// SetDomain - Set the Atlassian tenant domain name (e.g. samcaldwell in samcaldwell.atlassian.net)
func (client *Client) SetDomain(domain *string) error {
	if domain == nil || *domain == "" {
		return fmt.Errorf("missing domain (use --domain <value>)")
	}
	return client.domain.Set(domain)
}
