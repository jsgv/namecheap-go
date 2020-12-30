package api

import "encoding/xml"

// DomainsNSCreateResponse represents the response
// for the `namecheap.domains.ns.create` method.
type DomainsNSCreateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName              xml.Name `xml:"CommandResponse" json:"-"`
		Type                 string   `xml:"Type,attr"`
		DomainNSCreateResult struct {
			XMLName     xml.Name `xml:"DomainNSCreateResult" json:"-"`
			Domain      string   `xml:"Domain,attr"`
			Nameserver  string   `xml:"Nameserver,attr"`
			IP          string   `xml:"IP,attr"`
			IsSuccess   bool     `xml:"IsSuccess,attr"`
			ErrorNo     int      `xml:"ErrorNo,attr" json:",omitempty"`
			Description string   `xml:"Description" json:",omitempty"`
		}
	}
}

// DomainsNSCreateOptions represents the options
// form the `namecheap.domains.ns.create` method.
type DomainsNSCreateOptions struct {
	SLD        string
	TLD        string
	Nameserver string
	IP         string
}

// DomainsNSCreate creates a new nameserver.
func (c *Client) DomainsNSCreate(opts DomainsNSCreateOptions) (*DomainsNSCreateResponse, error) {
	r := &DomainsNSCreateResponse{}
	err := c.do("namecheap.domains.ns.create", opts, r)
	return r, err
}

// DomainsNSDeleteResponse represents the response
// for the `namecheap.domains.ns.delete` method.
type DomainsNSDeleteResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName              xml.Name `xml:"CommandResponse" json:"-"`
		Type                 string   `xml:"Type,attr"`
		DomainNSDeleteResult struct {
			XMLName     xml.Name `xml:"DomainNSDeleteResult" json:"-"`
			Domain      string   `xml:"Domain,attr"`
			Nameserver  string   `xml:"Nameserver,attr"`
			IsSuccess   bool     `xml:"IsSuccess,attr"`
			ErrorNo     int      `xml:"ErrorNo,attr" json:",omitempty"`
			Description string   `xml:"Description,attr" json:",omitempty"`
		}
	}
}

// DomainsNSDeleteOptions represents the options
// form the `namecheap.domains.ns.delete` method.
type DomainsNSDeleteOptions struct {
	SLD        string
	TLD        string
	Nameserver string
}

// DomainsNSDelete deletes a nameserver.
func (c *Client) DomainsNSDelete(opts DomainsNSDeleteOptions) (*DomainsNSDeleteResponse, error) {
	r := &DomainsNSDeleteResponse{}
	err := c.do("namecheap.domains.ns.delete", opts, r)
	return r, err
}

// DomainsNSGetInfoResponse represents the response
// for the `namecheap.domains.ns.getInfo` method.
type DomainsNSGetInfoResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName            xml.Name `xml:"CommandResponse" json:"-"`
		Type               string   `xml:"Type,attr"`
		DomainNSInfoResult struct {
			XMLName     xml.Name `xml:"DomainNSInfoResult" json:"-"`
			Domain      string   `xml:"Domain,attr"`
			Nameserver  string   `xml:"Nameserver,attr" json:",omitempty"`
			IP          string   `xml:"IP,attr" json:",omitempty"`
			ErrorNo     int      `xml:"ErrorNo,attr" json:",omitempty"`
			Description string   `xml:"Description,attr" json:",omitempty"`
		}
	}
}

// DomainsNSGetInfoOptions represents the options
// form the `namecheap.domains.ns.getInfo` method.
type DomainsNSGetInfoOptions struct {
	SLD        string
	TLD        string
	Nameserver string
}

// DomainsNSGetInfo retrieves information about a registered nameserver.
func (c *Client) DomainsNSGetInfo(opts DomainsNSGetInfoOptions) (*DomainsNSGetInfoResponse, error) {
	r := &DomainsNSGetInfoResponse{}
	err := c.do("namecheap.domains.ns.getInfo", opts, r)
	return r, err
}

// DomainsNSUpdateResponse represents the response
// for the `namecheap.domains.ns.update` method.
type DomainsNSUpdateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName              xml.Name `xml:"CommandResponse" json:"-"`
		Type                 string   `xml:"Type,attr"`
		DomainNSUpdateResult struct {
			XMLName     xml.Name `xml:"DomainNSUpdateResult" json:"-"`
			Domain      string   `xml:"Domain,attr"`
			Nameserver  string   `xml:"Nameserver,attr" json:",omitempty"`
			IsSuccess   bool     `xml:"IsSuccess,attr"`
			ErrorNo     int      `xml:"ErrorNo,attr" json:",omitempty"`
			Description string   `xml:"Description,attr" json:",omitempty"`
		}
	}
}

// DomainsNSUpdateOptions represents the options
// form the `namecheap.domains.ns.update` method.
type DomainsNSUpdateOptions struct {
	SLD        string
	TLD        string
	Nameserver string
	OldIP      string
	IP         string
}

// DomainsNSUpdate updates the IP address of a registered nameserver.
func (c *Client) DomainsNSUpdate(opts DomainsNSUpdateOptions) (*DomainsNSUpdateResponse, error) {
	r := &DomainsNSUpdateResponse{}
	err := c.do("namecheap.domains.ns.update", opts, r)
	return r, err
}
