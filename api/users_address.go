package api

import "encoding/xml"

// UsersAddressCreateResponse represents the response
// for the `namecheap.users.address.create` method.
type UsersAddressCreateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		AddressCreateResult struct {
			XMLName     xml.Name `xml:"AddressCreateResult" json:"-"`
			Success     bool     `xml:"Success,attr"`
			AddressId   int      `xml:"AddressId,attr"`
			AddressName string   `xml:"AddressName,attr"`
		}
	}
}

// UsersAddressCreateOptions represents the options
// for the `namecheap.users.address.create` method.
type UsersAddressCreateOptions struct {
	AddressName         string
	DefaultYN           int
	EmailAddress        string
	FirstName           string
	LastName            string
	JobTitle            string
	Organization        string
	Address1            string
	Address2            string
	City                string
	StateProvince       string
	StateProvinceChoice string
	Zip                 string
	Country             string
	Phone               string
	PhoneExt            string
	Fax                 string
}

// UsersAddressCreate creates a new address for the user.
func (c *Client) UsersAddressCreate(opts UsersAddressCreateOptions) (*UsersAddressCreateResponse, error) {
	r := &UsersAddressCreateResponse{}
	err := c.do("namecheap.users.address.create", opts, r)
	return r, err
}

// UsersAddressDeleteResponse represents the response
// for the `namecheap.users.address.delete` method.
type UsersAddressDeleteResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		AddressDeleteResult struct {
			XMLName   xml.Name `xml:"AddressDeleteResult" json:"-"`
			Success   bool     `xml:"Success,attr"`
			ProfileId int      `xml:"ProfileId,attr"`
			UserName  string   `xml:"UserName,attr"`
		}
	}
}

// UsersAddressDeleteOptions represents the options
// for the `namecheap.users.address.delete` method.
type UsersAddressDeleteOptions struct {
	AddressId int
}

// UsersAddressDelete deletes the particular address for the user.
func (c *Client) UsersAddressDelete(opts UsersAddressDeleteOptions) (*UsersAddressDeleteResponse, error) {
	r := &UsersAddressDeleteResponse{}
	err := c.do("namecheap.users.address.delete", opts, r)
	return r, err
}

// UsersAddressGetInfoResponse represents the response
// for the `namecheap.users.address.getInfo` method.
type UsersAddressGetInfoResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		GetAddressInfoResult struct {
			XMLName             xml.Name `xml:"GetAddressInfoResult" json:"-"`
			AddressId           int      `xml:"AddressId"`
			UserName            string   `xml:"UserName"`
			AddressName         string   `xml:"AddressName"`
			Default_YN          bool     `xml:"Default_YN"`
			FirstName           string   `xml:"FirstName"`
			LastName            string   `xml:"LastName"`
			JobTitle            string   `xml:"JobTitle"`
			Organization        string   `xml:"Organization"`
			Address1            string   `xml:"Address1"`
			Address2            string   `xml:"Address2"`
			City                string   `xml:"City"`
			StateProvince       string   `xml:"StateProvince"`
			StateProvinceChoice string   `xml:"StateProvinceChoice"`
			Zip                 string   `xml:"Zip"`
			Country             string   `xml:"Country"`
			Phone               string   `xml:"Phone"`
			PhoneExt            string   `xml:"PhoneExt"`
			EmailAddress        string   `xml:"EmailAddress"`
		}
	}
}

// UsersAddressGetInfoOptions represents the options
// for the `namecheap.users.address.getInfo` method.
type UsersAddressGetInfoOptions struct {
	AddressId int
}

// UsersAddressGetInfo gets information for the requested addressID.
func (c *Client) UsersAddressGetInfo(opts UsersAddressGetInfoOptions) (*UsersAddressGetInfoResponse, error) {
	r := &UsersAddressGetInfoResponse{}
	err := c.do("namecheap.users.address.getInfo", opts, r)
	return r, err
}

// UsersAddressGetListResponse represents the response
// for the `namecheap.users.address.getList` method.
type UsersAddressGetListResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		AddressGetListResult struct {
			XMLName xml.Name `xml:"AddressGetListResult" json:"-"`
			List    []struct {
				XMLName     xml.Name `xml:"List" json:"-"`
				AddressId   int      `xml:"AddressId,attr"`
				AddressName string   `xml:"AddressName,attr"`
				IsDefault   bool     `xml:"IsDefault,attr"`
			}
		}
	}
}

// UsersAddressGetList gets a list of addressIDs and addressnames associated with the user account.
func (c *Client) UsersAddressGetList() (*UsersAddressGetListResponse, error) {
	r := &UsersAddressGetListResponse{}
	err := c.do("namecheap.users.address.getList", nil, r)
	return r, err
}

// UsersAddressSetDefaultResponse represents the response
// for the `namecheap.users.address.setDefault` method.
type UsersAddressSetDefaultResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		AddressSetDefaultResult struct {
			XMLName   xml.Name `xml:"AddressSetDefaultResult" json:"-"`
			Success   bool     `xml:"Success,attr"`
			AddressId int      `xml:"AddressId,attr"`
		}
	}
}

// UsersAddressSetDefaultOptions represents the options
// for the `namecheap.users.address.setDefault` method.
type UsersAddressSetDefaultOptions struct {
	AddressId int
}

// UsersAddressSetDefault sets default address for the user.
func (c *Client) UsersAddressSetDefault(opts UsersAddressSetDefaultOptions) (*UsersAddressSetDefaultResponse, error) {
	r := &UsersAddressSetDefaultResponse{}
	err := c.do("namecheap.users.address.setDefault", opts, r)
	return r, err
}

// UsersAddressUpdateResponse represents the response
// for the `namecheap.users.address.update` method.
type UsersAddressUpdateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		AddressUpdateResult struct {
			XMLName     xml.Name `xml:"AddressUpdateResult" json:"-"`
			Success     bool     `xml:"Success,attr"`
			AddressId   int      `xml:"AddressId,attr"`
			AddressName string   `xml:"AddressName,attr"`
		}
	}
}

// UsersAddressUpdateOptions represents the options
// for the `namecheap.users.address.update` method.
type UsersAddressUpdateOptions struct {
	AddressId           int
	AddressName         string
	DefaultYN           int
	EmailAddress        string
	FirstName           string
	LastName            string
	JobTitle            string
	Organization        string
	Address1            string
	Address2            string
	City                string
	StateProvince       string
	StateProvinceChoice string
	Zip                 string
	Country             string
	Phone               string
	PhoneExt            string
	Fax                 string
}

// UsersAddressUpdate updates the particular address of the user.
func (c *Client) UsersAddressUpdate(opts UsersAddressUpdateOptions) (*UsersAddressUpdateResponse, error) {
	r := &UsersAddressUpdateResponse{}
	err := c.do("namecheap.users.address.update", opts, r)
	return r, err
}
