package api

import "encoding/xml"

// WhoisguardChangeEmailAddressResponse represents the response
// for the `namecheap.whoisguard.changeemailaddress` method.
type WhoisguardChangeEmailAddressResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                            xml.Name `xml:"CommandResponse" json:"-"`
		Type                               string   `xml:"Type,attr"`
		WhoisguardChangeEmailAddressResult struct {
			XMLName    xml.Name `xml:"WhoisguardChangeEmailAddressResult" json:"-"`
			ID         int      `xml:"ID,attr"`
			IsSuccess  bool     `xml:"IsSuccess,attr"`
			WGEmail    string   `xml:"WGEmail,attr"`
			WGOldEmail string   `xml:"WGOldEmail,attr"`
		}
	}
}

// WhoisguardChangeEmailAddressOptions represents the options
// for the `namecheap.whoisguard.changeemailaddress` method.
type WhoisguardChangeEmailAddressOptions struct {
	WhoisguardID int
}

// WhoisguardChangeEmailAddress changes domain privacy email address.
func (c *Client) WhoisguardChangeEmailAddress(opts WhoisguardChangeEmailAddressOptions) (*WhoisguardChangeEmailAddressResponse, error) {
	r := &WhoisguardChangeEmailAddressResponse{}
	err := c.do("namecheap.whoisguard.changeemailaddress", opts, r)
	return r, err
}

// WhoisguardEnableResponse represents the response
// for the `namecheap.whoisguard.enable` method.
type WhoisguardEnableResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                xml.Name `xml:"CommandResponse" json:"-"`
		Type                   string   `xml:"Type,attr"`
		WhoisguardEnableResult struct {
			XMLName    xml.Name `xml:"WhoisguardEnableResult" json:"-"`
			DomainName int      `xml:"DomainName,attr"`
			IsSuccess  bool     `xml:"IsSuccess,attr"`
		}
	}
}

// WhoisguardEnableOptions represents the options
// for the `namecheap.whoisguard.enable` method.
type WhoisguardEnableOptions struct {
	WhoisguardID     int
	ForwardedToEmail string
}

// WhoisguardEnable enables domain privacy protection for the WhoisguardID.
func (c *Client) WhoisguardEnable(opts WhoisguardEnableOptions) (*WhoisguardEnableResponse, error) {
	r := &WhoisguardEnableResponse{}
	err := c.do("namecheap.whoisguard.enable", opts, r)
	return r, err
}

// WhoisguardDisableResponse represents the response
// for the `namecheap.whoisguard.disable` method.
type WhoisguardDisableResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		WhoisguardDisableResult struct {
			XMLName    xml.Name `xml:"WhoisguardDisableResult" json:"-"`
			DomainName int      `xml:"DomainName,attr"`
			IsSuccess  bool     `xml:"IsSuccess,attr"`
		}
	}
}

// WhoisguardDisableOptions represents the options
// for the `namecheap.whoisguard.disable` method.
type WhoisguardDisableOptions struct {
	WhoisguardID int
}

// WhoisguardDisable disables domain privacy protection for the WhoisguardID.
func (c *Client) WhoisguardDisable(opts WhoisguardDisableOptions) (*WhoisguardDisableResponse, error) {
	r := &WhoisguardDisableResponse{}
	err := c.do("namecheap.whoisguard.disable", opts, r)
	return r, err
}

// WhoisguardGetListResponse represents the response
// for the `namecheap.whoisguard.getList` method.
type WhoisguardGetListResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		WhoisguardGetListResult struct {
			XMLName    xml.Name `xml:"WhoisguardGetListResult" json:"-"`
			Whoisguard []struct {
				XMLName    xml.Name `xml:"Whoisguard" json:"-"`
				ID         int      `xml:"ID,attr"`
				DomainName string   `xml:"DomainName,attr"`
				Created    string   `xml:"Created,attr"`
				Expires    string   `xml:"Expires,attr"`
				Status     string   `xml:"Status,attr"`
			}
		}
	}
}

// WhoisguardGetListOptions represents the options
// for the `namecheap.whoisguard.getList` method.
type WhoisguardGetListOptions struct {
	ListType string
	Page     int
	PageSize int
}

// WhoisguardGetList gets the list of domain privacy protection.
func (c *Client) WhoisguardGetList(opts WhoisguardGetListOptions) (*WhoisguardGetListResponse, error) {
	r := &WhoisguardGetListResponse{}
	err := c.do("namecheap.whoisguard.getList", opts, r)
	return r, err
}

// WhoisguardRenewResponse represents the response
// for the `namecheap.whoisguard.renew` method.
type WhoisguardRenewResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName               xml.Name `xml:"CommandResponse" json:"-"`
		Type                  string   `xml:"Type,attr"`
		WhoisguardRenewResult struct {
			XMLName       xml.Name `xml:"WhoisguardRenewResult" json:"-"`
			WhoisguardId  int      `xml:"WhoisguardId,attr"`
			Years         int      `xml:"Years,attr"`
			Renew         bool     `xml:"Renew,attr"`
			OrderId       int      `xml:"OrderId,attr"`
			TransactionId int      `xml:"TransactionId,attr"`
			ChargedAmount string   `xml:"ChargedAmount,attr"`
		}
	}
}

// WhoisguardRenewOptions represents the options
// for the `namecheap.whoisguard.renew` method.
type WhoisguardRenewOptions struct {
	WhoisguardID  string
	Years         int
	PromotionCode int
}

// WhoisguardRenew renews domain privacy protection.
func (c *Client) WhoisguardRenew(opts WhoisguardRenewOptions) (*WhoisguardRenewResponse, error) {
	r := &WhoisguardRenewResponse{}
	err := c.do("namecheap.whoisguard.renew", opts, r)
	return r, err
}
