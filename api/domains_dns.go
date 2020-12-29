package api

import "encoding/xml"

// DomainsDnsSetCustomResponse represents the response
// for the `namecheap.domains.dns.setCustom` method.
type DomainsDnsSetCustomResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                  xml.Name `xml:"CommandResponse" json:"-"`
		Type                     string   `xml:"Type,attr"`
		DomainDNSSetCustomResult struct {
			XMLName xml.Name `xml:"DomainDNSSetCustomResult" json:"-"`
			Domain  string   `xml:"Domain,attr"`
			Updated bool     `xml:"Updated,attr"`
		}
	}
}

// DomainsDnsSetCustomOptions represents the options
// for the `namecheap.domains.getList` method.
type DomainsDnsSetCustomOptions struct {
	SLD         string
	TLD         string
	Nameservers string
}

// DomainsDnsSetCustom sets domain to use custom DNS servers.
func (c *Client) DomainsDnsSetCustom(opts DomainsDnsSetCustomOptions) (*DomainsDnsSetCustomResponse, error) {
	r := &DomainsDnsSetCustomResponse{}
	err := c.do("namecheap.domains.dns.setCustom", opts, r)
	return r, err
}
