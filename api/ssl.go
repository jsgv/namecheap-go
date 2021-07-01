package api

import "encoding/xml"

// SSLCreateResponse represents the response
// for the `namecheap.ssl.create` method.
type SSLCreateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName         xml.Name `xml:"CommandResponse" json:"-"`
		Type            string   `xml:"Type,attr"`
		SSLCreateResult struct {
			XMLName        xml.Name `xml:"SSLCreateResult" json:"-"`
			IsSuccess      bool     `xml:"IsSuccess,attr"`
			OrderID        int      `xml:"OrderId,attr"`
			TransactionID  int      `xml:"TransactionId,attr"`
			ChargedAmount  string   `xml:"ChargedAmount,attr"`
			SSLCertificate struct {
				XMLName       xml.Name `xml:"SSLCertificate" json:"-"`
				CertificateID int      `xml:"CertificateID,attr" json:",omitempty"`
				Created       string   `xml:"Created,attr" json:",omitempty"`
				SSLType       string   `xml:"SSLType,attr" json:",omitempty"`
				Years         int      `xml:"Years,attr" json:",omitempty"`
				Status        string   `xml:"Status,attr" json:",omitempty"`
			}
		}
	}
}

// SSLCreateOptions represents the options
// for the `namecheap.ssl.create` method.
type SSLCreateOptions struct {
	Years         int
	Type          string
	SANStoADD     int `namecheap:"intnozero"`
	PromotionCode string
}

// SSLCreate creates a new SSL certificate.
func (c *Client) SSLCreate(opts SSLCreateOptions) (*SSLCreateResponse, error) {
	r := &SSLCreateResponse{}
	err := c.do("namecheap.ssl.create", opts, r)
	return r, err
}

// SSLGetListResponse represents the response
// for the `namecheap.ssl.getList` method.
type SSLGetListResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName       xml.Name `xml:"CommandResponse" json:"-"`
		Type          string   `xml:"Type,attr"`
		Paging        Paging
		SSLListResult struct {
			XMLName xml.Name `xml:"SSLListResult" json:"-"`
			SSLs    []struct {
				CertificateID        int    `xml:"CertificateID,attr"`
				HostName             string `xml:"HostName,attr" json:",omitempty"`
				SSLType              string `xml:"SSLType,attr"`
				PurchaseDate         string `xml:"PurchaseDate,attr"`
				ExpireDate           string `xml:"ExpireDate,attr" json:",omitempty"`
				ActivationExpireDate string `xml:"ActivationExpireDate,attr" json:",omitempty"`
				IsExpiredYN          bool   `xml:"IsExpiredYN,attr"`
				Status               string `xml:"Status,attr"`
				ProviderOrderID      int    `xml:"ProviderOrderID,attr" json:",omitempty"`
				Years                int    `xml:"Years,attr"`
			} `xml:"SSL"`
		}
	}
}

// SSLGetListOptions represents the options
// for the `namecheap.ssl.getList` method.
type SSLGetListOptions struct {
	ListType   string
	SearchTerm string
	Page       int `namecheap:"intnozero"`
	PageSize   int `namecheap:"intnozero"`
	SortBy     string
}

// SSLGetList returns a list of SSL certificates for a particular user.
func (c *Client) SSLGetList(opts SSLGetListOptions) (*SSLGetListResponse, error) {
	r := &SSLGetListResponse{}
	err := c.do("namecheap.ssl.getList", opts, r)
	return r, err
}

// SSLParseCSRResponse represents the response
// for the `namecheap.ssl.parseCSR` method.
type SSLParseCSRResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName           xml.Name `xml:"CommandResponse" json:"-"`
		Type              string   `xml:"Type,attr"`
		SSLParseCSRResult struct {
			XMLName    xml.Name `xml:"SSLParseCSRResult" json:"-"`
			CSRDetails struct {
				CommonName       string `xml:",chardata"`
				DomainName       string `xml:",chardata"`
				Country          string `xml:",chardata"`
				OrganisationUnit string `xml:",chardata"`
				Organisation     string `xml:",chardata"`
				ValidTrueDomain  bool   `xml:",chardata"`
				State            string `xml:",chardata"`
				Locality         string `xml:",chardata"`
				Email            string `xml:",chardata"`
				DNSNames         string `xml:",chardata"`
			}
		}
	}
}

// SSLParseCSROptions represents the options
// for the `namecheap.ssl.parseCSR` method.
type SSLParseCSROptions struct {
	CSR             string
	CertificateType string
}

// SSLParseCSR parses the CSR.
func (c *Client) SSLParseCSR(opts SSLParseCSROptions) (*SSLParseCSRResponse, error) {
	r := &SSLParseCSRResponse{}
	err := c.do("namecheap.ssl.parseCSR", opts, r)
	return r, err
}

// SSLGetApproverEmailListResponse represents the response
// for the `namecheap.ssl.getApproverEmailList` method.
type SSLGetApproverEmailListResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                    xml.Name `xml:"CommandResponse" json:"-"`
		Type                       string   `xml:"Type,attr"`
		GetApproverEmailListResult struct {
			XMLName      xml.Name `xml:"GetApproverEmailListResult" json:"-"`
			Domain       string   `xml:"Domain,attr"`
			Domainemails struct {
				Emails []string `xml:"email"`
			}
			Genericemails struct {
				Emails []string `xml:"email"`
			}
		}
	}
}

// SSLGetApproverEmailListOptions represents the options
// for the `namecheap.ssl.getApproverEmailList` method.
type SSLGetApproverEmailListOptions struct {
	DomainName      string
	CertificateType string
}

// SSLGetApproverEmailList gets approver email list for the requested domain.
func (c *Client) SSLGetApproverEmailList(opts SSLGetApproverEmailListOptions) (*SSLGetApproverEmailListResponse, error) {
	r := &SSLGetApproverEmailListResponse{}
	err := c.do("namecheap.ssl.getApproverEmailList", opts, r)
	return r, err
}

// SSLActivateResponse represents the response
// for the `namecheap.ssl.activate` method.
type SSLActivateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName           xml.Name `xml:"CommandResponse" json:"-"`
		Type              string   `xml:"Type,attr"`
		SSLActivateResult struct {
			XMLName   xml.Name `xml:"SSLActivateResult" json:"-"`
			ID        int      `xml:"ID,attr"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
		}
	}
}

// SSLActivateOptions represents the options
// for the `namecheap.ssl.activate` method.
type SSLActivateOptions struct {
	CertificateID     string
	CSR               string
	AdminEmailAddress string
	WebServerType     string

	// Domain Control Validation (DCV)
	ApproverEmail    string
	HTTPDCValidation *bool
	DNSDCValidation  *bool

	// Multi-domain specific
	DNSNames          string
	DNSApproverEmails string

	// OV & EV
	AdminOrganizationName  string
	OrganizationDepartment string
	AdminCountry           string
	AdminStateProvince     string
	AdminCity              string
	AdminAddress1          string
	AdminAddress2          string
	AdminPostalCode        string
	AdminPhone             string
	OrganizationDUNS       string

	// EV only
	CompanyIncorporationCountry       string
	CompanyIncorporationStateProvince string
	CompanyIncorporationLocality      string
	CompanyIncorporationDate          string
	CompanyDBA                        string
	CompanyRegistrationNumber         string

	// OV only
	OrganizationRepFirstName    string
	OrganizationRepLastName     string
	OrganizationRepTitle        string
	OrganizationRepPhone        string
	OrganizationRepEmailAddress string
}

// SSLActivate gets approver email list for the requested domain.
func (c *Client) SSLActivate(opts SSLActivateOptions) (*SSLActivateResponse, error) {
	r := &SSLActivateResponse{}
	err := c.do("namecheap.ssl.activate", opts, r)
	return r, err
}

// SSLResendApproverEmailResponse represents the response
// for the `namecheap.ssl.resendApproverEmail` method.
type SSLResendApproverEmailResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                      xml.Name `xml:"CommandResponse" json:"-"`
		Type                         string   `xml:"Type,attr"`
		SSLResendApproverEmailResult struct {
			XMLName   xml.Name `xml:"SSLResendApproverEmailResult" json:"-"`
			ID        int      `xml:"ID,attr"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
		}
	}
}

// SSLResendApproverEmailOptions represents the options
// for the `namecheap.ssl.resendApproverEmail` method.
type SSLResendApproverEmailOptions struct {
	CertificateID string
}

// SSLResendApproverEmail resends the approver email.
func (c *Client) SSLResendApproverEmail(opts SSLResendApproverEmailOptions) (*SSLResendApproverEmailResponse, error) {
	r := &SSLResendApproverEmailResponse{}
	err := c.do("namecheap.ssl.resendApproverEmail", opts, r)
	return r, err
}

// SSLGetInfoResponse represents the response
// for the `namecheap.ssl.getInfo` method.
type SSLGetInfoResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName          xml.Name `xml:"CommandResponse" json:"-"`
		Type             string   `xml:"Type,attr"`
		SSLGetInfoResult struct {
			XMLName              xml.Name `xml:"SSLGetInfoResult" json:"-"`
			Status               string   `xml:"Status,attr"`
			StatusDescription    string   `xml:"StatusDescription,attr"`
			Type                 string   `xml:"Type,attr"`
			IssuedOn             string   `xml:"IssuedOn,attr"`
			Years                int      `xml:"Years,attr"`
			Expires              string   `xml:"Expires,attr"`
			ActivationExpireDate string   `xml:"ActivationExpireDate,attr"`
			OrderId              int      `xml:"OrderId,attr"`
			SANSCount            int      `xml:"SANSCount,attr"`
			Provider             struct {
				XMLName xml.Name `xml:"Provider" json:"-"`
				Name    string   `xml:"Name"`
			}
			// TODO:
			// CertificateDetails struct
		}
	}
}

// SSLGetInfoOptions represents the options
// for the `namecheap.ssl.getInfo` method.
type SSLGetInfoOptions struct {
	CertificateID     string
	Returncertificate string
	Returntype        string
}

// SSLGetInfo retrieves information about the requested SSL certificate.
func (c *Client) SSLGetInfo(opts SSLGetInfoOptions) (*SSLGetInfoResponse, error) {
	r := &SSLGetInfoResponse{}
	err := c.do("namecheap.ssl.getInfo", opts, r)
	return r, err
}

// SSLRenewResponse represents the response
// for the `namecheap.ssl.renew` method.
type SSLRenewResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName        xml.Name `xml:"CommandResponse" json:"-"`
		Type           string   `xml:"Type,attr"`
		SSLRenewResult struct {
			XMLName       xml.Name `xml:"SSLRenewResult" json:"-"`
			CertificateID int      `xml:"CertificateID,attr"`
			SSLType       string   `xml:"SSLType,attr"`
			Years         int      `xml:"Years,attr"`
			OrderID       int      `xml:"OrderID,attr"`
			TransactionID int      `xml:"TransactionID,attr"`
			ChargedAmount string   `xml:"ChargedAmount,attr"`
		}
	}
}

// SSLRenewOptions represents the options
// for the `namecheap.ssl.renew` method.
type SSLRenewOptions struct {
	CertificateID string
	Years         int
	SSLType       string
	PromotionCode string
}

// SSLRenew renews an SSL certificate.
func (c *Client) SSLRenew(opts SSLRenewOptions) (*SSLRenewResponse, error) {
	r := &SSLRenewResponse{}
	err := c.do("namecheap.ssl.renew", opts, r)
	return r, err
}

// SSLReissueResponse represents the response
// for the `namecheap.ssl.reissue` method.
type SSLReissueResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName xml.Name `xml:"CommandResponse" json:"-"`
		Type    string   `xml:"Type,attr"`

		// NOTE: Docs show `SSLActivateResult` node
		// But I have only seen `SSLReissueResult` node

		// TODO: finish rest of nodes
		SSLReissueResult struct {
			XMLName   xml.Name `xml:"SSLReissueResult" json:"-"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
			ID        int      `xml:"ID,attr"`
		}
	}
}

// SSLReissueOptions represents the options
// for the `namecheap.ssl.reissue` method.
type SSLReissueOptions struct {
	CertificateID     string
	CSR               string
	AdminEmailAddress string
	WebServerType     string

	// Domain Control Validation (DCV)
	ApproverEmail    string
	HTTPDCValidation *bool
	DNSDCValidation  *bool

	// Multi-domain specific
	DNSNames          string
	DNSApproverEmails string

	// OV & EV
	AdminOrganizationName  string
	OrganizationDepartment string
	AdminCountry           string
	AdminStateProvince     string
	AdminCity              string
	AdminAddress1          string
	AdminAddress2          string
	AdminPostalCode        string
	AdminPhone             string
	OrganizationDUNS       string

	// EV only
	CompanyIncorporationCountry       string
	CompanyIncorporationStateProvince string
	CompanyIncorporationLocality      string
	CompanyIncorporationDate          string
	CompanyDBA                        string
	CompanyRegistrationNumber         string

	// OV only
	OrganizationRepFirstName    string
	OrganizationRepLastName     string
	OrganizationRepTitle        string
	OrganizationRepPhone        string
	OrganizationRepEmailAddress string
}

// SSLReissue reissues an SSL certificate.
func (c *Client) SSLReissue(opts SSLReissueOptions) (*SSLReissueResponse, error) {
	r := &SSLReissueResponse{}
	err := c.do("namecheap.ssl.reissue", opts, r)
	return r, err
}

// SSLResendFulfillmentEmailResponse represents the response
// for the `namecheap.ssl.resendfulfillmentemail` method.
type SSLResendFulfillmentEmailResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                         xml.Name `xml:"CommandResponse" json:"-"`
		Type                            string   `xml:"Type,attr"`
		SSLResendFulfillmentEmailResult struct {
			XMLName   xml.Name `xml:"SSLResendFulfillmentEmailResult" json:"-"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
			ID        int      `xml:"ID,attr"`
		}
	}
}

// SSLResendFulfillmentEmailOptions represents the options
// for the `namecheap.ssl.resendfulfillmentemail` method.
type SSLResendFulfillmentEmailOptions struct {
	CertificateID string
}

// SSLResendFulfillmentEmail resends the fulfilment email containing the certificate.
func (c *Client) SSLResendFulfillmentEmail(opts SSLResendFulfillmentEmailOptions) (*SSLResendFulfillmentEmailResponse, error) {
	r := &SSLResendFulfillmentEmailResponse{}
	err := c.do("namecheap.ssl.resendfulfillmentemail", opts, r)
	return r, err
}

// SSLPurchaseMoreSansResponse represents the response
// for the `namecheap.ssl.purchasemoresans` method.
type SSLPurchaseMoreSansResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                   xml.Name `xml:"CommandResponse" json:"-"`
		Type                      string   `xml:"Type,attr"`
		SSLPurchaseMoreSANSResult struct {
			XMLName       xml.Name `xml:"SSLPurchaseMoreSANSResult" json:"-"`
			IsSuccess     bool     `xml:"IsSuccess,attr"`
			OrderID       int      `xml:"OrderId,attr"`
			TransactionID int      `xml:"TransactionId,attr"`
			ChargedAmount string   `xml:"ChargedAmount,attr"`

			SSLCertificate struct {
				XMLName       xml.Name `xml:"SSLCertificate" json:"-"`
				CertificateID int      `xml:"CertificateID,attr"`
				SSLType       string   `xml:"SSLType,attr"`
				Years         int      `xml:"Years,attr"`
				Status        string   `xml:"Status,attr"`
				SANSCount     int      `xml:"SANSCount,attr"`
			}
		}
	}
}

// SSLPurchaseMoreSansOptions represents the options
// for the `namecheap.ssl.purchasemoresans` method.
type SSLPurchaseMoreSansOptions struct {
	CertificateID     string
	NumberOfSANSToAdd int
}

// SSLPurchaseMoreSans purchases more add-on domains for an already purchased certificate.
func (c *Client) SSLPurchaseMoreSans(opts SSLPurchaseMoreSansOptions) (*SSLPurchaseMoreSansResponse, error) {
	r := &SSLPurchaseMoreSansResponse{}
	err := c.do("namecheap.ssl.purchasemoresans", opts, r)
	return r, err
}

// SSLRevokeCertificateResponse represents the response
// for the `namecheap.ssl.revokecertificate` method.
type SSLRevokeCertificateResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                 xml.Name `xml:"CommandResponse" json:"-"`
		Type                    string   `xml:"Type,attr"`
		RevokeCertificateResult struct {
			XMLName   xml.Name `xml:"RevokeCertificateResult" json:"-"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
			ID        int      `xml:"ID,attr"`
		}
	}
}

// SSLRevokeCertificateOptions represents the options
// for the `namecheap.ssl.revokecertificate` method.
type SSLRevokeCertificateOptions struct {
	CertificateID   int
	CertificateType string
}

// SSLRevokeCertificate revokes a re-issued SSL certificate.
func (c *Client) SSLRevokeCertificate(opts SSLRevokeCertificateOptions) (*SSLRevokeCertificateResponse, error) {
	r := &SSLRevokeCertificateResponse{}
	err := c.do("namecheap.ssl.revokecertificate", opts, r)
	return r, err
}

// SSLEditDCVMethodResponse represents the response
// for the `namecheap.ssl.editDCVMethod` method.
type SSLEditDCVMethodResponse struct {
	ApiResponse

	CommandResponse struct {
		XMLName                xml.Name `xml:"CommandResponse" json:"-"`
		Type                   string   `xml:"Type,attr"`
		SSLEditDCVMethodResult struct {
			XMLName   xml.Name `xml:"SSLEditDCVMethodResult" json:"-"`
			IsSuccess bool     `xml:"IsSuccess,attr"`
			ID        int      `xml:"ID,attr"`

			HttpDCValidation struct {
				XMLName        xml.Name `xml:"HttpDCValidation" json:"-"`
				ValueAvailable bool     `xml:"ValueAvailable,attr" json:",omitempty"`
			}

			// TODO: missing nodes
		}
	}
}

// SSLEditDCVMethodOptions represents the options
// for the `namecheap.ssl.editDCVMethod` method.
type SSLEditDCVMethodOptions struct {
	CertificateID int
	DCVMethod     string
	DNSNames      string
	DCVMethods    string
}

// SSLEditDCVMethod sets new domain control validation (DCV) method for a certificate or serves as 'retry' mechanism.
func (c *Client) SSLEditDCVMethod(opts SSLEditDCVMethodOptions) (*SSLEditDCVMethodResponse, error) {
	r := &SSLEditDCVMethodResponse{}
	err := c.do("namecheap.ssl.editDCVMethod", opts, r)
	return r, err
}
