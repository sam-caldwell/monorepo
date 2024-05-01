package Client

func (client *Client) SetDomain(domain string) error {
    return client.domain.Set(&domain)
}
