package api

// DomainsDnsSetCustomResponse represents the response
// for the `namecheap.domains.dns.setCustom` method.
type DomainsDnsSetCustomResponse struct {
	ApiResponse
}

// DomainsDnsSetCustomOptions represents the options
// for the `namecheap.domains.getList` method.
type DomainsDnsSetCustomOptions struct {
	SLD         string
	TLD         string
	Nameservers string
}

// DomainsDnsSetCustom sets domain to use custom DNS servers.
func (c *Client) DomainsDnsSetCustom(opts DomainsDnsSetCustomOptions) *DomainsDnsSetCustomResponse {
	r := &DomainsDnsSetCustomResponse{}
	c.Do("namecheap.domains.dns.setCustom", opts, r)
	return r
}
