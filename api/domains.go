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
	Page       int
	PageSize   int
	SortBy     string
}

// DomainsGetList returns a list of domains for the particular user.
func (c *Client) DomainsGetList(opts DomainsGetListOptions) (*DomainsGetListResponse, error) {
	r := &DomainsGetListResponse{}
	err := c.do("namecheap.domains.getList", opts, r)
	return r, err
}

type contactInfo struct {
	OrganizationName    string `xml:"OrganizationName" json:",omitempty"`
	JobTitle            string `xml:"JobTitle" json:",omitempty"`
	FirstName           string `xml:"FirstName" json:",omitempty"`
	LastName            string `xml:"LastName" json:",omitempty"`
	Address1            string `xml:"Address1" json:",omitempty"`
	Address2            string `xml:"Address2" json:",omitempty"`
	City                string `xml:"City" json:",omitempty"`
	StateProvince       string `xml:"StateProvince" json:",omitempty"`
	StateProvinceChoice string `xml:"StateProvinceChoice" json:",omitempty"`
	PostalCode          string `xml:"PostalCode" json:",omitempty"`
	Country             string `xml:"Country" json:",omitempty"`
	Phone               string `xml:"Phone" json:",omitempty"`
	Fax                 string `xml:"Fax" json:",omitempty"`
	EmailAddress        string `xml:"EmailAddress" json:",omitempty"`
	PhoneExt            string `xml:"PhoneExt" json:",omitempty"`
	FaxExt              string `xml:"FaxExt" json:",omitempty"`
}

type contactList struct {
	Registrant struct {
		XMLName  xml.Name `xml:"Registrant" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr" json:",omitempty"`
		contactInfo
	} `json:",omitempty"`
	Tech struct {
		XMLName  xml.Name `xml:"Tech" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr" json:",omitempty"`
		contactInfo
	} `json:",omitempty"`
	Admin struct {
		XMLName  xml.Name `xml:"Admin" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr" json:",omitempty"`
		contactInfo
	} `json:",omitempty"`
	AuxBilling struct {
		XMLName  xml.Name `xml:"AuxBilling" json:"-"`
		ReadOnly bool     `xml:"ReadOnly,attr" json:",omitempty"`
		contactInfo
	} `json:",omitempty"`
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
			contactList
			// CurrentAttributes // TODO
			WhoisGuardContact struct {
				XMLName xml.Name `xml:"WhoisGuardContact" json:"-"`
				contactList
			} `json:",omitempty"`
		}
	}
}

// DomainsGetContactsOptions represents the options
// for the `namecheap.domains.getContacts` method.
type DomainsGetContactsOptions struct {
	DomainName string
}

// DomainsGetContacts returns contact information for the requested domain.
func (c *Client) DomainsGetContacts(opts DomainsGetContactsOptions) (*DomainsGetContactsResponse, error) {
	r := &DomainsGetContactsResponse{}
	err := c.do("namecheap.domains.getContacts", opts, r)
	return r, err
}

// DomainsCreateResponse represents the response
// for the `namecheap.domains.create` method.
type DomainsCreateResponse struct {
	ApiResponse
	CommandResponse struct {
		XMLName            xml.Name `xml:"CommandResponse" json:"-"`
		Type               string   `xml:"Type,attr"`
		DomainCreateResult struct {
			XMLName           xml.Name `xml:"DomainCreateResult" json:"-"`
			Registered        bool     `xml:"Registered,attr"`
			ChargedAmount     string   `xml:"ChargedAmount,attr"`
			DomainID          int      `xml:"DomainID,attr"`
			OrderID           int      `xml:"OrderID,attr"`
			TransactionID     int      `xml:"TransactionID,attr"`
			WhoisguardEnable  bool     `xml:"WhoisguardEnable,attr"`
			FreePositiveSSL   bool     `xml:"FreePositiveSSL,attr"`
			NonRealTimeDomain bool     `xml:"NonRealTimeDomain,attr"`
		}
	}
}

// DomainsCreateOptions represents the options
// for the `namecheap.domains.create` method.
type DomainsCreateOptions struct {
	DomainName                    string
	Years                         string
	PromotionCode                 string
	BillingFirstName              string
	BillingLastName               string
	BillingAddress1               string
	BillingAddress2               string
	BillingCity                   string
	BillingStateProvince          string
	BillingStateProvinceChoice    string
	BillingPostalCode             string
	BillingCountry                string
	BillingPhone                  string
	BillingPhoneExt               string
	BillingFax                    string
	BillingEmailAddress           string
	IdnCode                       string
	Nameservers                   string
	AddFreeWhoisguard             string
	WGEnabled                     string
	IsPremiumDomain               bool
	PremiumPrice                  string
	EapFee                        string
	RegistrantOrganizationName    string
	RegistrantJobTitle            string
	RegistrantFirstName           string
	RegistrantLastName            string
	RegistrantAddress1            string
	RegistrantAddress2            string
	RegistrantCity                string
	RegistrantStateProvince       string
	RegistrantStateProvinceChoice string
	RegistrantPostalCode          string
	RegistrantCountry             string
	RegistrantPhone               string
	RegistrantPhoneExt            string
	RegistrantFax                 string
	RegistrantEmailAddress        string
	TechOrganizationName          string
	TechJobTitle                  string
	TechFirstName                 string
	TechLastName                  string
	TechAddress1                  string
	TechAddress2                  string
	TechCity                      string
	TechStateProvince             string
	TechStateProvinceChoice       string
	TechPostalCode                string
	TechCountry                   string
	TechPhone                     string
	TechPhoneExt                  string
	TechFax                       string
	TechEmailAddress              string
	AdminOrganizationName         string
	AdminJobTitle                 string
	AdminFirstName                string
	AdminLastName                 string
	AdminAddress1                 string
	AdminAddress2                 string
	AdminCity                     string
	AdminStateProvince            string
	AdminStateProvinceChoice      string
	AdminPostalCode               string
	AdminCountry                  string
	AdminPhone                    string
	AdminPhoneExt                 string
	AdminFax                      string
	AdminEmailAddress             string
	AuxBillingOrganizationName    string
	AuxBillingJobTitle            string
	AuxBillingFirstName           string
	AuxBillingLastName            string
	AuxBillingAddress1            string
	AuxBillingAddress2            string
	AuxBillingCity                string
	AuxBillingStateProvince       string
	AuxBillingStateProvinceChoice string
	AuxBillingPostalCode          string
	AuxBillingCountry             string
	AuxBillingPhone               string
	AuxBillingPhoneExt            string
	AuxBillingFax                 string
	AuxBillingEmailAddress        string
	// ExtendedAttributes ?
}

func (c *Client) DomainsCreate(opts DomainsCreateOptions) (*DomainsCreateResponse, error) {
	r := &DomainsCreateResponse{}
	err := c.do("namecheap.domains.create", opts, r)
	return r, err
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
func (c *Client) DomainsGetTldList() (*DomainsGetTldListResponse, error) {
	var opts interface{}
	r := &DomainsGetTldListResponse{}
	err := c.do("namecheap.domains.getTldList", opts, r)
	return r, err
}

// DomainsSetContactsResponse represents the response
// for the `namecheap.domains.setContacts` method.
type DomainsSetContactsResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                xml.Name `xml:"CommandResponse" json:"-"`
		Type                   string   `xml:"Type,attr"`
		DomainSetContactResult struct {
			XMLName   xml.Name `xml:"DomainSetContactResult" json:"-"`
			Domain    string   `xml:"Domain,attr"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
			Warnings  struct {
				XMLName xml.Name `xml:"Warnings" json:"-"`
				Warning struct {
					XMLName xml.Name `xml:"Warning" json:"-"`
					Number  int      `xml:"Number,attr" json:",omitempty"`
					Warning string   `xml:",chardata" json:",omitempty"`
				}
			}
		}
	}
}

// DomainsSetContactsOptions represents the options
// for the `namecheap.domains.setContacts` method.
type DomainsSetContactsOptions struct {
	DomainName                    string
	RegistrantOrganizationName    string
	RegistrantJobTitle            string
	RegistrantFirstName           string
	RegistrantLastName            string
	RegistrantAddress1            string
	RegistrantAddress2            string
	RegistrantCity                string
	RegistrantStateProvince       string
	RegistrantStateProvinceChoice string
	RegistrantPostalCode          string
	RegistrantCountry             string
	RegistrantPhone               string
	RegistrantPhoneExt            string
	RegistrantFax                 string
	RegistrantEmailAddress        string
	TechOrganizationName          string
	TechJobTitle                  string
	TechFirstName                 string
	TechLastName                  string
	TechAddress1                  string
	TechAddress2                  string
	TechCity                      string
	TechStateProvince             string
	TechStateProvinceChoice       string
	TechPostalCode                string
	TechCountry                   string
	TechPhone                     string
	TechPhoneExt                  string
	TechFax                       string
	TechEmailAddress              string
	AdminOrganizationName         string
	AdminJobTitle                 string
	AdminFirstName                string
	AdminLastName                 string
	AdminAddress1                 string
	AdminAddress2                 string
	AdminCity                     string
	AdminStateProvince            string
	AdminStateProvinceChoice      string
	AdminPostalCode               string
	AdminCountry                  string
	AdminPhone                    string
	AdminPhoneExt                 string
	AdminFax                      string
	AdminEmailAddress             string
	AuxBillingOrganizationName    string
	AuxBillingJobTitle            string
	AuxBillingFirstName           string
	AuxBillingLastName            string
	AuxBillingAddress1            string
	AuxBillingAddress2            string
	AuxBillingCity                string
	AuxBillingStateProvince       string
	AuxBillingStateProvinceChoice string
	AuxBillingPostalCode          string
	AuxBillingCountry             string
	AuxBillingPhone               string
	AuxBillingPhoneExt            string
	AuxBillingFax                 string
	AuxBillingEmailAddress        string
	// Extended attributes ?
}

// DomainsSetContacts sets contact information for the requested domain.
func (c *Client) DomainsSetContacts(opts DomainsSetContactsOptions) (*DomainsSetContactsResponse, error) {
	r := &DomainsSetContactsResponse{}
	err := c.do("namecheap.domains.setContacts", opts, r)
	return r, err
}

// DomainsCheckResponse represents the response
// for the `namecheap.domains.check` method.
type DomainsCheckResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName           xml.Name `xml:"CommandResponse" json:"-"`
		Type              string   `xml:"Type,attr"`
		DomainCheckResult struct {
			XMLName                  xml.Name `xml:"DomainCheckResult" json:"-"`
			Domain                   string   `xml:"Domain,attr"`
			Available                bool     `xml:"Available,attr"`
			ErrorNo                  int      `xml:"ErrorNo,attr"`
			Description              string   `xml:"Description,attr"`
			IsPremiumName            bool     `xml:"IsPremiumName,attr"`
			PremiumRegistrationPrice string   `xml:"PremiumRegistrationPrice,attr"`
			PremiumRenewalPrice      string   `xml:"PremiumRenewalPrice,attr"`
			PremiumRestorePrice      string   `xml:"PremiumRestorePrice,attr"`
			PremiumTransferPrice     string   `xml:"PremiumTransferPrice,attr"`
			IcannFee                 string   `xml:"IcannFee,attr"`
			EapFee                   string   `xml:"EapFee,attr"`
		}
	}
}

// DomainsCheckOptions represents the options
// for the `namecheap.domains.check` method.
type DomainsCheckOptions struct {
	DomainList string
}

// DomainsCheck checks the availability of domains.
func (c *Client) DomainsCheck(opts DomainsCheckOptions) (*DomainsCheckResponse, error) {
	r := &DomainsCheckResponse{}
	err := c.do("namecheap.domains.check", opts, r)
	return r, err
}

// DomainsReactivateResponse represents the response
// from the `namecheap.domains.reactivate` method.
type DomainsReactivateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`
	}
}

// DomainsReactivateOptions represents the options
// for the `namecheap.domains.reactivate` method.
type DomainsReactivateOptions struct {
	DomainName      string
	PromotionCode   string
	YearsToAdd      int
	IsPremiumDomain bool
	PremiumPrice    string
}

// DomainsReactivate reactivates an expired domain.
func (c *Client) DomainsReactivate(opts DomainsReactivateOptions) (*DomainsReactivateResponse, error) {
	r := &DomainsReactivateResponse{}
	err := c.do("namecheap.domains.reactivate", opts, r)
	return r, err
}

// DomainRenewResponse represents the response
// for the `namecheap.domains.renew` method.
type DomainRenewResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName           xml.Name `xml:"CommandResponse" json:"-"`
		Type              string   `xml:"Type,attr"`
		DomainRenewResult struct {
			XMLName       xml.Name `xml:"DomainRenewResult" json:"-"`
			DomainName    string   `xml:"DomainName,attr"`
			DomainID      int      `xml:"DomainID,attr"`
			Renew         bool     `xml:"Renew,attr"`
			OrderID       int      `xml:"OrderID,attr"`
			TransactionID int      `xml:"TransactionID,attr"`
			ChargedAmount string   `xml:"ChargedAmount,attr"`
			DomainDetails struct {
				XMLName     xml.Name `xml:"DomainDetails" json:"-"`
				ExpiredDate string   `xml:"ExpiredDate"`
				NumYears    int      `xml:"NumYears"`
			}
		}
	}
}

// DomainRenewOptions represents the options
// for the `namecheap.domains.renew` method.
type DomainRenewOptions struct {
	DomainName      string
	Years           int
	PromotionCode   string
	IsPremiumDomain bool
	PremiumPrice    string
}

// DomainsRenew renews an expiring domain.
func (c *Client) DomainsRenew(opts DomainRenewOptions) (*DomainRenewResponse, error) {
	r := &DomainRenewResponse{}
	err := c.do("namecheap.domains.renew", opts, r)
	return r, err
}

// DomainsGetRegistrarLockResponse represents the response
// for the `namecheap.domains.getRegistrarLock` method.
type DomainsGetRegistrarLockResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                      xml.Name `xml:"CommandResponse" json:"-"`
		Type                         string   `xml:"Type,attr"`
		DomainGetRegistrarLockResult struct {
			XMLName                  xml.Name `xml:"DomainGetRegistrarLockResult" json:"-"`
			RegistrarLockStatus      bool     `xml:"RegistrarLockStatus,attr"`
			IsClientUpdateProhibited bool     `xml:"IsClientUpdateProhibited,attr"`
			IsClientDeleteProhibited bool     `xml:"IsClientDeleteProhibited,attr"`
			IsClientHold             bool     `xml:"IsClientHold,attr"`
		}
	}
}

// DomainsGetRegistrarLockOptions represents the options
// for the `namecheap.domains.getRegistrarLock` method.
type DomainsGetRegistrarLockOptions struct {
	DomainName string
}

// DomainsGetRegistrarLock returns the Registrar Lock status for the requested domain.
func (c *Client) DomainsGetRegistrarLock(opts DomainsGetRegistrarLockOptions) (*DomainsGetRegistrarLockResponse, error) {
	r := &DomainsGetRegistrarLockResponse{}
	err := c.do("namecheap.domains.getRegistrarLock", opts, r)
	return r, err
}

// DomainsSetRegistrarLockResponse represents the response
// for the `namecheap.domains.setRegistrarLock` method.
type DomainsSetRegistrarLockResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                      xml.Name `xml:"CommandResponse" json:"-"`
		Type                         string   `xml:"Type,attr"`
		DomainSetRegistrarLockResult struct {
			XMLName                      xml.Name `xml:"DomainSetRegistrarLockResult" json:"-"`
			Domain                       string   `xml:"Domain,attr"`
			IsSuccess                    bool     `xml:"IsSuccess,attr"`
			RegistrarLockStatus          bool     `xml:"RegistrarLockStatus,attr"`
			IsRegistrarLockStatusUpdated bool     `xml:"IsRegistrarLockStatusUpdated,attr"`
			IsClientUpdateProhibited     bool     `xml:"IsClientUpdateProhibited,attr"`
			IsClientDeleteProhibited     bool     `xml:"IsClientDeleteProhibited,attr"`
			IsClientHoldUpdated          bool     `xml:"IsClientHoldUpdated,attr"`
		}
	}
}

// DomainsSetRegistrarLockOptions represents the options
// for the `namecheap.domains.setRegistrarLock` method.
type DomainsSetRegistrarLockOptions struct {
	DomainName string
	LockAction string
}

// DomainsSetRegistrarLock sets the Registrar Lock status for a domain
func (c *Client) DomainsSetRegistrarLock(opts DomainsSetRegistrarLockOptions) (*DomainsSetRegistrarLockResponse, error) {
	r := &DomainsSetRegistrarLockResponse{}
	err := c.do("namecheap.domains.setRegistrarLock", opts, r)
	return r, err
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
func (c *Client) DomainsGetInfo(opts DomainsGetInfoOptions) (*DomainsGetInfoResponse, error) {
	r := &DomainsGetInfoResponse{}
	err := c.do("namecheap.domains.getInfo", opts, r)
	return r, err
}
