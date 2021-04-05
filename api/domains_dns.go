package api

import (
	"encoding/xml"
	"fmt"
)

// DomainsDNSSetDefaultResponse represents the response
// for the `namecheap.domains.dns.setDefault` method.
type DomainsDNSSetDefaultResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                   xml.Name `xml:"CommandResponse" json:"-"`
		Type                      string   `xml:"Type,attr"`
		DomainDNSSetDefaultResult struct {
			XMLName xml.Name `xml:"DomainDNSSetDefaultResult" json:"-"`
			Domain  string   `xml:"Domain,attr"`
			Updated bool     `xml:"Updated,attr"`
		}
	}
}

// DomainsDNSSetDefaultOptions represent the options
// for the `namecheap.domains.dns.setDefault` method.
type DomainsDNSSetDefaultOptions struct {
	SLD string
	TLD string
}

// DomainsDNSSetDefault sets domain to use Namecheap default DNS servers.
func (c *Client) DomainsDNSSetDefault(opts DomainsDNSSetDefaultOptions) (*DomainsDNSSetDefaultResponse, error) {
	r := &DomainsDNSSetDefaultResponse{}
	err := c.do("namecheap.domains.dns.setDefault", opts, r)
	return r, err
}

// DomainsDNSSetCustomResponse represents the response
// for the `namecheap.domains.dns.setCustom` method.
type DomainsDNSSetCustomResponse struct {
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

// DomainsDNSSetCustomOptions represents the options
// for the `namecheap.domains.dns.setCustom` method.
type DomainsDNSSetCustomOptions struct {
	SLD         string
	TLD         string
	Nameservers string
}

// DomainsDNSSetCustom sets domain to use custom DNS servers.
func (c *Client) DomainsDNSSetCustom(opts DomainsDNSSetCustomOptions) (*DomainsDNSSetCustomResponse, error) {
	r := &DomainsDNSSetCustomResponse{}
	err := c.do("namecheap.domains.dns.setCustom", opts, r)
	return r, err
}

// DomainsDNSGetListResponse represents the response
// for the `namecheap.domains.dns.getList` method.
type DomainsDNSGetListResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                xml.Name `xml:"CommandResponse" json:"-"`
		Type                   string   `xml:"Type,attr"`
		DomainDNSGetListResult struct {
			XMLName        xml.Name `xml:"DomainDNSGetListResult" json:"-"`
			Domain         string   `xml:"Domain,attr"`
			IsUsingOurDNS  bool     `xml:"IsUsingOurDNS,attr"`
			IsPremiumDNS   bool     `xml:"IsPremiumDNS,attr"`
			IsUsingFreeDNS bool     `xml:"IsUsingFreeDNS,attr"`
			Nameserver     []struct {
				Nameserver string `xml:",chardata"`
			}
		}
	}
}

// DomainsDNSGetListOptions represents the options
// for the `namecheap.domains.dns.getList` method.
type DomainsDNSGetListOptions struct {
	SLD string
	TLD string
}

// DomainsDNSGetList gets a list of DNS servers associated with the requested domain.
func (c *Client) DomainsDNSGetList(opts DomainsDNSGetListOptions) (*DomainsDNSGetListResponse, error) {
	r := &DomainsDNSGetListResponse{}
	err := c.do("namecheap.domains.dns.getList", opts, r)
	return r, err
}

// DomainsDNSGetHostsResponse represents the response
// for the `namecheap.domains.dns.getHosts` method.
type DomainsDNSGetHostsResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		DomainDNSGetHostsResult struct {
			XMLName       xml.Name `xml:"DomainDNSGetHostsResult" json:"-"`
			Domain        string   `xml:"Domain,attr"`
			EmailType     string   `xml:"EmailType,attr"`
			IsUsingOurDNS bool     `xml:"IsUsingOurDNS,attr"`
			Hosts         []struct {
				HostId             string `xml:"HostId,attr"`
				Name               string `xml:"Name,attr"`
				Type               string `xml:"Type,attr"`
				Address            string `xml:"Address,attr"`
				MXPref             int    `xml:"MXPref,attr"`
				TTL                int    `xml:"TTL,attr"`
				AssociatedAppTitle string `xml:"AssociatedAppTitle,attr"`
				FriendlyName       string `xml:"FriendlyName,attr"`
				IsActive           bool   `xml:"IsActive,attr"`
				IsDDNSEnabled      bool   `xml:"IsDDNSEnabled,attr"`
			} `xml:"host"`
		}
	}
}

// DomainsDNSGetHostsOptions represents the options
// for the `namecheap.domains.dns.getHosts` method.
type DomainsDNSGetHostsOptions struct {
	SLD string
	TLD string
}

// DomainsDNSGetHosts retrieves DNS host record settings for the requested domain.
func (c *Client) DomainsDNSGetHosts(opts DomainsDNSGetHostsOptions) (*DomainsDNSGetHostsResponse, error) {
	r := &DomainsDNSGetHostsResponse{}
	err := c.do("namecheap.domains.dns.getHosts", opts, r)
	return r, err
}

// DomainsDNSGetEmailForwardingResponse represents the response
// for the `namecheap.domains.dns.getEmailForwarding` method.
type DomainsDNSGetEmailForwardingResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                           xml.Name `xml:"CommandResponse" json:"-"`
		Type                              string   `xml:"Type,attr"`
		DomainDNSGetEmailForwardingResult struct {
			XMLName xml.Name `xml:"DomainDNSGetEmailForwardingResult" json:"-"`
			Domain  string   `xml:"Domain,attr"`
		}
	}
}

// DomainsDNSGetEmailForwardingOptions represents the options
// for the `namecheap.domains.dns.getEmailForwarding` method.
type DomainsDNSGetEmailForwardingOptions struct {
	DomainName string
}

// DomainsDNSGetEmailForwarding gets email forwarding settings for the requested domain.
func (c *Client) DomainsDNSGetEmailForwarding(opts DomainsDNSGetEmailForwardingOptions) (*DomainsDNSGetEmailForwardingResponse, error) {
	r := &DomainsDNSGetEmailForwardingResponse{}
	err := c.do("namecheap.domains.dns.getEmailForwarding", opts, r)
	return r, err
}

// DomainsDNSSetEmailForwardingResponse represents the response
// for the `namecheap.domains.dns.setEmailForwarding` method.
type DomainsDNSSetEmailForwardingResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                           xml.Name `xml:"CommandResponse" json:"-"`
		Type                              string   `xml:"Type,attr"`
		DomainDNSSetEmailForwardingResult struct {
			XMLName   xml.Name `xml:"DomainDNSSetEmailForwardingResult" json:"-"`
			Domain    string   `xml:"Domain,attr"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
		}
	}
}

// DomainsDNSSetEmailForwardingOptions represents the options
// for the `namecheap.domains.dns.setEmailForwarding` method.
type DomainsDNSSetEmailForwardingOptions struct {
	DomainName string
	MailBox    []string
	ForwardTo  []string
}

// DomainsDNSSetEmailForwarding sets email forwarding for a domain name.
func (c *Client) DomainsDNSSetEmailForwarding(opts DomainsDNSSetEmailForwardingOptions) (*DomainsDNSSetEmailForwardingResponse, error) {
	m := map[string]interface{}{
		"DomainName": opts.DomainName,
	}

	for i, v := range opts.MailBox {
		k := fmt.Sprintf("MailBox%d", i+1)
		m[k] = v
	}

	for i, v := range opts.ForwardTo {
		k := fmt.Sprintf("ForwardTo%d", i+1)
		m[k] = v
	}

	r := &DomainsDNSSetEmailForwardingResponse{}
	err := c.do("namecheap.domains.dns.setEmailForwarding", m, r)
	return r, err
}

// DomainsDNSSetHostsResponse represents the response
// for the `namecheap.domains.dns.setHosts` method.
type DomainsDNSSetHostsResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		DomainDNSSetHostsResult struct {
			XMLName   xml.Name `xml:"DomainDNSSetHostsResult" json:"-"`
			Domain    string   `xml:"Domain,attr"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
			EmailType string   `xml:"EmailType,attr"`
		}
	}
}

// DomainsDNSSetHostsOptions represents the options
// for the `namecheap.domains.dns.setHosts` method.
type DomainsDNSSetHostsOptions struct {
	SLD        string
	TLD        string
	EmailType  string
	Flag       string
	Tag        string
	HostName   []string
	RecordType []string
	Address    []string
	MXPref     []string
	TTL        []string
}

// DomainsDNSSetHosts sets DNS host records settings for the requested domain.
func (c *Client) DomainsDNSSetHosts(opts DomainsDNSSetHostsOptions) (*DomainsDNSSetHostsResponse, error) {
	m := map[string]interface{}{
		"SLD":       opts.SLD,
		"TLD":       opts.TLD,
		"EmailType": opts.EmailType,
		"Flag":      opts.Flag,
		"Tag":       opts.Tag,
	}

	for i, v := range opts.HostName {
		k := fmt.Sprintf("HostName%d", i+1)
		m[k] = v
	}

	for i, v := range opts.RecordType {
		k := fmt.Sprintf("RecordType%d", i+1)
		m[k] = v
	}

	for i, v := range opts.Address {
		k := fmt.Sprintf("Address%d", i+1)
		m[k] = v
	}

	for i, v := range opts.MXPref {
		k := fmt.Sprintf("MXPref%d", i+1)
		m[k] = v
	}

	for i, v := range opts.TTL {
		k := fmt.Sprintf("TTL%d", i+1)
		m[k] = v
	}

	r := &DomainsDNSSetHostsResponse{}

	err := c.do("namecheap.domains.dns.setHosts", m, r)
	return r, err
}
