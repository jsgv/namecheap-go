package api

import (
	"encoding/xml"
)

// DomainsTransferCreateResponse represents the response
// for the `namecheap.domains.transfer.create` method.
type DomainsTransferCreateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                    xml.Name `xml:"CommandResponse" json:"-"`
		Type                       string   `xml:"Type,attr"`
		DomainTransferCreateResult struct {
			XMLName       xml.Name `xml:"DomainTransferCreateResult" json:"-"`
			DomainName    string   `xml:"DomainName,attr"`
			Transfer      bool     `xml:"Transfer,attr"`
			TransferID    int      `xml:"TransferID,attr"`
			StatusID      int      `xml:"StatusID,attr"`
			OrderID       int      `xml:"OrderID,attr"`
			TransactionID int      `xml:"TransactionID,attr"`
			ChargedAmount string   `xml:"ChargedAmount,attr"`
			StatusCode    int      `xml:"StatusCode,attr"`
		}
	}
}

// DomainsTransferCreateOptions represents the options
// for the `namecheap.domains.transfer.create` method.
type DomainsTransferCreateOptions struct {
	DomainName        string
	Years             string
	EPPCode           string
	PromotionCode     string
	AddFreeWhoisguard bool `namecheap:"yesno"`
	WGenable          bool `namecheap:"yesno"`
}

// DomainsTransferCreate transfers a domain to Namecheap.
func (c *Client) DomainsTransferCreate(opts DomainsTransferCreateOptions) (*DomainsTransferCreateResponse, error) {
	r := &DomainsTransferCreateResponse{}
	err := c.do("namecheap.domains.transfer.create", opts, r)
	return r, err
}

// DomainsTransferGetStatusResponse represents the response
// for the `namecheap.domains.transfer.getStatus` method.
type DomainsTransferGetStatusResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                       xml.Name `xml:"CommandResponse" json:"-"`
		Type                          string   `xml:"Type,attr"`
		DomainTransferGetStatusResult struct {
			XMLName           xml.Name `xml:"DomainTransferGetStatusResult" json:"-"`
			TransferID        int      `xml:"TransferID,attr"`
			Status            string   `xml:"Status,attr"`
			StatusID          int      `xml:"StatusID,attr"`
			StatusDate        string   `xml:"StatusDate,attr"`
			TransferOrderDate string   `xml:"TransferOrderDate,attr"`
		}
	}
}

// DomainsTransferGetStatusOptions represents the options
// for the `namecheap.domains.transfer.getStatus` method.
type DomainsTransferGetStatusOptions struct {
	TransferID int
}

// DomainsTransferGetStatus gets the status of a particular transfer.
func (c *Client) DomainsTransferGetStatus(opts DomainsTransferGetStatusOptions) (*DomainsTransferGetStatusResponse, error) {
	r := &DomainsTransferGetStatusResponse{}
	err := c.do("namecheap.domains.transfer.getStatus", opts, r)
	return r, err
}

// DomainsTransferUpdateStatusResponse represents the response
// for the `namecheap.domains.transfer.updateStatus` method.
type DomainsTransferUpdateStatusResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                          xml.Name `xml:"CommandResponse" json:"-"`
		Type                             string   `xml:"Type,attr"`
		DomainTransferUpdateStatusResult struct {
			XMLName    xml.Name `xml:"DomainTransferUpdateStatusResult" json:"-"`
			TransferID int      `xml:"TransferId,attr"`
			Resubmit   bool     `xml:"Resubmit,attr"`
		}
	}
}

// DomainsTransferUpdateStatusOptions represents the options
// for the `namecheap.domains.transfer.updateStatus` method.
type DomainsTransferUpdateStatusOptions struct {
	TransferID int
	Resubmit   bool
}

// DomainsTransferUpdateStatus updates the status of a particular transfer.
func (c *Client) DomainsTransferUpdateStatus(opts DomainsTransferUpdateStatusOptions) (*DomainsTransferUpdateStatusResponse, error) {
	r := &DomainsTransferUpdateStatusResponse{}
	err := c.do("namecheap.domains.transfer.updateStatus", opts, r)
	return r, err
}

// DomainsTransferGetListResponse represents the response
// for the `namecheap.domains.transfer.getList` method.
type DomainsTransferGetListResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName               xml.Name `xml:"CommandResponse" json:"-"`
		Type                  string   `xml:"Type,attr"`
		Paging                Paging
		TransferGetListResult struct {
			XMLName   xml.Name `xml:"TransferGetListResult" json:"-"`
			Transfers []struct {
				ID                int    `xml:"ID,attr"`
				DomainName        string `xml:"DomainName,attr"`
				User              string `xml:"User,attr"`
				TransferDate      string `xml:"TransferDate,attr"`
				OrderID           int    `xml:"OrderID,attr"`
				StatusID          int    `xml:"StatusID,attr"`
				Status            string `xml:"Status,attr"`
				StatusDate        string `xml:"StatusDate,attr"`
				StatusDescription string `xml:"StatusDescription,attr"`
			} `xml:"Transfer"`
		}
	}
}

// DomainsTransferGetListOptions represents the options
// for the `namecheap.domains.transfer.getList` method.
type DomainsTransferGetListOptions struct {
	ListType   string
	SearchTerm string
	Page       int
	PageSize   int
	SortBy     string
}

// DomainsTransferGetList gets the list of domain transfers.
func (c *Client) DomainsTransferGetList(opts DomainsTransferGetListOptions) (*DomainsTransferGetListResponse, error) {
	r := &DomainsTransferGetListResponse{}
	err := c.do("namecheap.domains.transfer.getList", opts, r)
	return r, err
}
