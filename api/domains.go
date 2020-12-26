package api

import (
	"encoding/xml"
)

// DomainsGetListResponse represents the response
// for the `namecheap.domains.getList` method.
type DomainsGetListResponse struct {
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
func (c *Client) DomainsGetList(opts DomainsGetListOptions) *DomainsGetListResponse {
	r := &DomainsGetListResponse{}
	c.Do("namecheap.domains.getList", opts, r)
	return r
}

type ContactInfo struct {
	OrganizationName    string `xml:"OrganizationName"`
	JobTitle            string `xml:"JobTitle"`
	FirstName           string `xml:"FirstName"`
	LastName            string `xml:"LastName"`
	Address1            string `xml:"Address1"`
	Address2            string `xml:"Address2"`
	City                string `xml:"City"`
	StateProvince       string `xml:"StateProvince"`
	StateProvinceChoice string `xml:"StateProvinceChoice"`
	PostalCode          string `xml:"PostalCode"`
	Country             string `xml:"Country"`
	Phone               string `xml:"Phone"`
	Fax                 string `xml:"Fax"`
	EmailAddress        string `xml:"EmailAddress"`
	PhoneExt            string `xml:"PhoneExt"`
	FaxExt              string `xml:"FaxExt"`
}

type ContactList struct {
	Registrant struct {
		XMLName  xml.Name `xml:"Registrant" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr"`
		ContactInfo
	}
	Tech struct {
		XMLName  xml.Name `xml:"Tech" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr"`
		ContactInfo
	}
	Admin struct {
		XMLName  xml.Name `xml:"Admin" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr"`
		ContactInfo
	}
	AuxBilling struct {
		XMLName  xml.Name `xml:"AuxBilling" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr"`
		ContactInfo
	}
}

// DomainsGetContactsResponse represents the response
// for the `namecheap.domains.getContacts` method.
type DomainsGetContactsResponse struct {
	ApiResponse
	CommandResponse struct {
		XMLName              xml.Name `xml:"CommandResponse" json:"-"`
		Type                 string   `xml:"Type,attr"`
		DomainContactsResult struct {
			XMLName      xml.Name `xml:"DomainContactsResult" json:"-"`
			Domain       string   `xml:"Domain,attr"`
			Domainnameid string   `xml:"domainnameid,attr"`
			ContactList
			// CurrentAttributes // TODO
			WhoisGuardContact struct {
				XMLName xml.Name `xml:"WhoisGuardContact" json:"-"`
				ContactList
			}
		}
	}
}

// DomainsGetContactsOptions represents the options
// for the `namecheap.domains.getContacts` method.
type DomainsGetContactsOptions struct {
	DomainName string
}

// DomainsGetContacts returns contact information for the requested domain.
func (c *Client) DomainsGetContacts(opts DomainsGetContactsOptions) *DomainsGetContactsResponse {
	r := &DomainsGetContactsResponse{}
	c.Do("namecheap.domains.getContacts", opts, r)
	return r
}

// DomainsGetTldListResponse represents the response
// for the `namecheap.domains.getTldList` method.
type DomainsGetTldListResponse struct {
	ApiResponse
	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`
		Tlds    struct {
			XMLName xml.Name `xml:"Tlds" json:"-"`
			Tld     []struct {
				XMLName                       xml.Name `xml:"Tld" json:"-"`
				Value                         string   `xml:",chardata"`
				Name                          string   `xml:"Name,attr"`
				NonRealTime                   bool     `xml:"NonRealTime,attr"`
				MinRegisterYears              int      `xml:"MinRegisterYears,attr"`
				MaxRegisterYears              int      `xml:"MaxRegisterYears,attr"`
				MinRenewYears                 int      `xml:"MinRenewYears,attr"`
				MaxRenewYears                 int      `xml:"MaxRenewYears,attr"`
				RenewalMinDays                int      `xml:"RenewalMinDays,attr"`
				RenewalMaxDays                int      `xml:"RenewalMaxDays,attr"`
				ReactivateMaxDays             int      `xml:"ReactivateMaxDays,attr"`
				MinTransferYears              int      `xml:"MinTransferYears,attr"`
				MaxTransferYears              int      `xml:"MaxTransferYears,attr"`
				IsApiRegisterable             bool     `xml:"IsApiRegisterable,attr"`
				IsApiRenewable                bool     `xml:"IsApiRenewable,attr"`
				IsApiTransferable             bool     `xml:"IsApiTransferable,attr"`
				IsEppRequired                 bool     `xml:"IsEppRequired,attr"`
				IsDisableModContact           bool     `xml:"IsDisableModContact,attr"`
				IsDisableWGAllot              bool     `xml:"IsDisableWGAllot,attr"`
				IsIncludeInExtendedSearchOnly bool     `xml:"IsIncludeInExtendedSearchOnly,attr"`
				SequenceNumber                int      `xml:"SequenceNumber,attr"`
				Type                          string   `xml:"Type,attr"`
				SubType                       string   `xml:"SubType,attr"`
				IsSupportsIDN                 bool     `xml:"IsSupportsIDN,attr"`
				Category                      string   `xml:"Category,attr"`
				SupportsRegistrarLock         bool     `xml:"SupportsRegistrarLock,attr"`
				AddGracePeriodDays            int      `xml:"AddGracePeriodDays,attr"`
				WhoisVerification             bool     `xml:"WhoisVerification,attr"`
				ProviderApiDelete             bool     `xml:"ProviderApiDelete,attr"`
				TldState                      string   `xml:"TldState,attr"`
				SearchGroup                   string   `xml:"SearchGroup,attr"`
				Registry                      string   `xml:"Registry,attr"`
				Categories                    struct {
					XMLName     xml.Name `xml:"Categories" json:"-"`
					TldCategory struct {
						XMLName        xml.Name `xml:"TldCategory" json:"-"`
						Name           string   `xml:"Name,attr"`
						SequenceNumber string   `xml:"SequenceNumber,attr"`
					}
				}
			}
		}
	}
}

// DomainsGetTldList returns a list of TLDs
func (c *Client) DomainsGetTldList() *DomainsGetTldListResponse {
	var opts interface{}
	r := &DomainsGetTldListResponse{}
	c.Do("namecheap.domains.getTldList", opts, r)
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
