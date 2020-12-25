package api

import (
	"encoding/xml"
)

// DomainsGetListResultResponse represents the response
// for the `namecheap.domains.getList` method.
type DomainsGetListResultResponse struct {
	ApiResponse
	CommandResponse struct {
		XMLName             xml.Name `xml:"CommandResponse" json:"-"`
		Type                string   `xml:"Type,attr"`
		DomainGetListResult struct {
			XMLName xml.Name `xml:"DomainGetListResult" json:"-"`
			Domain  []struct {
				XMLName    xml.Name `xml:"Domain" json:"-"`
				ID         string   `xml:"ID,attr"`
				Name       string   `xml:"Name,attr"`
				User       string   `xml:"User,attr"`
				Created    string   `xml:"Created,attr"`
				Expires    string   `xml:"Expires,attr"`
				IsExpired  bool     `xml:"IsExpired,attr"`
				IsLocked   bool     `xml:"IsLocked,attr"`
				AutoRenew  bool     `xml:"AutoRenew,attr"`
				WhoisGuard string   `xml:"WhoisGuard,attr"`
				IsPremium  bool     `xml:"IsPremium,attr"`
				IsOurDNS   bool     `xml:"IsOurDNS,attr"`
			}
		}
		Paging struct {
			XMLName     xml.Name `xml:"Paging" json:"-"`
			TotalItems  int      `xml:"TotalItems"`
			CurrentPage int      `xml:"CurrentPage"`
			PageSize    int      `xml:"PageSize"`
		}
	}
}

// DomainsGetListOptions represents the options
// for the `namecheap.domains.getList` method.
type DomainsGetListOptions struct {
	ListType   string
	SearchTerm string
	Page       string
	PageSize   string
	SortBy     string
}

// DomainsGetList returns a list of domains for the particular user.
func (c *Client) DomainsGetList(opts DomainsGetListOptions) *DomainsGetListResultResponse {
	r := &DomainsGetListResultResponse{}
	c.Do("namecheap.domains.getList", opts, r)
	return r
}

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

// DomainsGetInfoResponse represents the response
// for the `namecheap.domains.getInfo` method.
type DomainsGetInfoResponse struct {
	ApiResponse
	CommandResponse struct {
		XMLName             xml.Name `xml:"CommandResponse" json:"-"`
		Type                string   `xml:"Type,attr"`
		DomainGetInfoResult struct {
			XMLName       xml.Name `xml:"DomainGetInfoResult" json:"-"`
			Status        string   `xml:"Status,attr"`
			ID            string   `xml:"ID,attr"`
			DomainName    string   `xml:"DomainName,attr"`
			OwnerName     string   `xml:"OwnerName,attr"`
			IsOwner       bool     `xml:"IsOwner,attr"`
			IsPremium     bool     `xml:"IsPremium,attr"`
			DomainDetails struct {
				XMLName     xml.Name `xml:"DomainDetails" json:"-"`
				CreatedDate string   `xml:"CreatedDate"`
				ExpiredDate string   `xml:"ExpiredDate"`
				NumYears    int      `xml:"NumYears"`
			}
			// TODO
			// LockDetails struct {
			// 	XMLName xml.Name `xml:"LockDetails" json:"-"`
			// }
			Whoisguard struct {
				XMLName      xml.Name `xml:"Whoisguard" json:"-"`
				Enabled      bool     `xml:"Enabled,attr"`
				ID           int      `xml:"ID"`
				ExpiredDate  string   `xml:"ExpiredDate"`
				EmailDetails struct {
					XMLName                      xml.Name `xml:"EmailDetails" json:"-"`
					WhoisGuardEmail              string   `xml:"WhoisGuardEmail,attr"`
					ForwardedTo                  string   `xml:"ForwardedTo,attr"`
					LastAutoEmailChangeDate      string   `xml:"LastAutoEmailChangeDate,attr"`
					AutoEmailChangeFrequencyDays int      `xml:"AutoEmailChangeFrequencyDays,attr"`
				}
			}
			PremiumDnsSubscription struct {
				XMLName        xml.Name `xml:"PremiumDnsSubscription" json:"-"`
				UseAutoRenew   bool     `xml:"UseAutoRenew"`
				SubscriptionId int      `xml:"SubscriptionId"`
				CreatedDate    string   `xml:"CreatedDate"`
				ExpirationDate string   `xml:"ExpirationDate"`
				IsActive       bool     `xml:"IsActive"`
			}
			DnsDetails struct {
				XMLName          xml.Name `xml:"DnsDetails" json:"-"`
				ProviderType     string   `xml:"ProviderType,attr"`
				IsUsingOurDNS    bool     `xml:"IsUsingOurDNS,attr"`
				HostCount        int      `xml:"HostCount,attr"`
				EmailType        string   `xml:"EmailType,attr"`
				DynamicDNSStatus bool     `xml:"DynamicDNSStatus,attr"`
				IsFailover       bool     `xml:"IsFailover,attr"`
				Nameserver       []string `xml:"Nameserver"`
			}
			Modificationrights struct {
				XMLName xml.Name `xml:"Modificationrights" json:"-"`
				All     bool     `xml:"All,attr"`
			}
		}
	}
}

// DomainsGetInfoOptions represents the options
// for the `namecheap.domains.getInfo` method.
type DomainsGetInfoOptions struct {
	DomainName string
	HostName   string
}

// DomainsGetInfo returns information about the requested domain.
func (c *Client) DomainsGetInfo(opts DomainsGetInfoOptions) *DomainsGetInfoResponse {
	r := &DomainsGetInfoResponse{}
	c.Do("namecheap.domains.getInfo", opts, r)
	return r
}
