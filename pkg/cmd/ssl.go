package cmd

import (
	"errors"
	"strconv"

	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdSSL() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ssl",
		Short: "Manage ssl",
	}

	addCommand(cmd, newCmdSSLCreate)
	addCommand(cmd, newCmdSSLGetList)
	addCommand(cmd, newCmdSSLParseCSR)
	addCommand(cmd, newCmdSSLGetApproverEmailList)
	addCommand(cmd, newCmdSSLActivate)
	addCommand(cmd, newCmdSSLResendApproverEmail)
	addCommand(cmd, newCmdSSLGetInfo)
	addCommand(cmd, newCmdSSLRenew)
	addCommand(cmd, newCmdSSLReissue)
	addCommand(cmd, newCmdSSLResendFulfillmentEmail)
	addCommand(cmd, newCmdSSLPurchaseMoreSans)
	addCommand(cmd, newCmdSSLRevokeCertificate)
	addCommand(cmd, newCmdSSLEditDCVMethod)

	return cmd, nil
}

func newCmdSSLCreate() (*cobra.Command, error) {
	var err error
	var opts api.SSLCreateOptions

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new SSL certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLCreate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.Years, "years", 0, "Number of years SSL will be issued for")
	err = cmd.MarkFlagRequired("years")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Type, "type", "", "SSL product name")
	err = cmd.MarkFlagRequired("type")
	if err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.SANStoADD, "sanstoadd", 0, "Defines number of add-on domains to be purchased in addition to the default number of domains included into a multi-domain certificate")
	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Promotional (coupon) code for the certificate")

	return cmd, nil
}

func newCmdSSLGetList() (*cobra.Command, error) {
	var opts api.SSLGetListOptions

	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Returns a list of SSL certificates for the particular user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLGetList(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.ListType, "listtype", "ALL", "Possible values are ALL, Processing, EmailSent, TechnicalProblem, InProgress, Completed, Deactivated, Active, Cancelled, NewPurchase, NewRenewal.")
	cmd.Flags().StringVar(&opts.SearchTerm, "searchterm", "", "Possible values are ALL, Processing, EmailSent, TechnicalProblem, InProgress, Completed, Deactivated, Active, Cancelled, NewPurchase, NewRenewal.")

	cmd.Flags().IntVar(&opts.Page, "page", 0, "Page to return")
	cmd.Flags().IntVar(&opts.PageSize, "pagesize", 0, "Total number of SSL certificates to display in a page. Minimum value is 10 and maximum value is 100")
	cmd.Flags().StringVar(&opts.SortBy, "sortby", "", "Possible values are PURCHASEDATE, PURCHASEDATE_DESC, SSLTYPE, SSLTYPE_DESC, EXPIREDATETIME, EXPIREDATETIME_DESC, Host_Name, Host_Name_DESC ")

	return cmd, nil
}

func newCmdSSLParseCSR() (*cobra.Command, error) {
	var (
		err  error
		opts api.SSLParseCSROptions
	)

	cmd := &cobra.Command{
		Use:   "parsecsr",
		Short: "Parses the CSR.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLParseCSR(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CSR, "csr", "", "Certificate Signing Request")
	err = cmd.MarkFlagRequired("csr")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.CertificateType, "certificatetype", "", "Type of SSL Certificate")

	return cmd, nil
}

func newCmdSSLGetApproverEmailList() (*cobra.Command, error) {
	var err error
	var opts api.SSLGetApproverEmailListOptions

	cmd := &cobra.Command{
		Use:   "getapproveremaillist",
		Short: "Gets approver email list for the requested certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLGetApproverEmailList(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to get the list")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.CertificateType, "certificatetype", "", "Type of SSL certificate")
	err = cmd.MarkFlagRequired("certificatetype")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdSSLActivate() (*cobra.Command, error) {
	var opts api.SSLActivateOptions
	var HTTPDCValidation, DNSDCValidation bool
	var err error

	cmd := &cobra.Command{
		Use:   "activate",
		Short: "Activates a newly purchased SSL certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpdcvFlag := cmd.Flag("httpdcvalidation")
			if httpdcvFlag.Changed {
				v, err := strconv.ParseBool(httpdcvFlag.Value.String())
				if err != nil {
					return err
				}
				opts.HTTPDCValidation = &v
			}

			dnsdcvFlag := cmd.Flag("dnsdcvalidation")
			if dnsdcvFlag.Changed {
				v, err := strconv.ParseBool(dnsdcvFlag.Value.String())
				if err != nil {
					return err
				}
				opts.DNSDCValidation = &v
			}

			r, err := apiClient.SSLActivate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "Unique identifier of SSL certificate to be activated")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.CSR, "csr", "", "Certificate Signing Request (CSR)")
	if err = cmd.MarkFlagRequired("csr"); err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.AdminEmailAddress, "adminemailaddress", "", "Email address to send signed SSL certificate file to")
	if err = cmd.MarkFlagRequired("adminemailaddress"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.ApproverEmail, "approveremail", "", "Sets method to pass DCV through")
	cmd.Flags().StringVar(&opts.WebServerType, "webservertype", "", "Server software where SSL will be installed")

	// Domain Control Validation (DCV)
	cmd.Flags().BoolVar(&HTTPDCValidation, "httpdcvalidation", false, "Sets all domains in certificate request to be validated through HTTP DCV")
	cmd.Flags().BoolVar(&DNSDCValidation, "dnsdcvalidation", false, "Sets all domains in certificate requested domains to be validated through CNAME DCV")

	// Multi-domain specific
	cmd.Flags().StringVar(&opts.DNSNames, "dnsnames", "", "Specifies domain names to be included into the certificate request in addition to the Common Name in (CSR) Format")
	cmd.Flags().StringVar(&opts.DNSApproverEmails, "dnsapproveremails", "", "Sets method to pass DCV through for each domain in the certificate request")

	// OV & EV
	cmd.Flags().StringVar(&opts.AdminOrganizationName, "adminorganizationname", "", "Organization name to be validated by Comodo validation staff and encoded into the issued certificate.")
	cmd.Flags().StringVar(&opts.OrganizationDepartment, "organizationdepartment", "", "Organizational Unit Name")
	cmd.Flags().StringVar(&opts.AdminCountry, "admincountry", "", "Organization country name to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminStateProvince, "adminstateprovince", "", "Organization state or province to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminCity, "admincity", "", "Organization city or locality to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminAddress1, "adminaddress1", "", "Organization street address to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminAddress2, "adminaddress2", "", "Additional organization street address to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminPostalCode, "adminpostalcode", "", "Organization postal code to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminPhone, "adminphone", "", "Organization phone number registered in any reputable online database used for certificate request verification callback")
	cmd.Flags().StringVar(&opts.OrganizationDUNS, "organizationduns", "", "Organization's number in DUN & Bradstreet company database")

	// EV only
	cmd.Flags().StringVar(&opts.CompanyIncorporationCountry, "companyincorporationcountry", "", "Country where the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyIncorporationStateProvince, "companyincorporationstateprovince", "", "State or province where the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyIncorporationLocality, "companyincorporationlocality", "", "City or locality where the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyIncorporationDate, "companyincorporationdate", "", "Date when the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyDBA, "companydba", "", "Organization's 'Doing business as' name")
	cmd.Flags().StringVar(&opts.CompanyRegistrationNumber, "companyregistrationnumber", "", "Number assigned to an organization by an Incorporating Agency")

	// OV only
	cmd.Flags().StringVar(&opts.OrganizationRepFirstName, "organizationrepfirstname", "", "Full-time organization representative's first name")
	cmd.Flags().StringVar(&opts.OrganizationRepLastName, "organizationreplastname", "", "Full-time organization representative's last name")
	cmd.Flags().StringVar(&opts.OrganizationRepTitle, "organizationreptitle", "", "Full-time organization representative's job title")
	cmd.Flags().StringVar(&opts.OrganizationRepPhone, "OrganizationRepPhone", "", "Full-time organization representative's phone number")
	cmd.Flags().StringVar(&opts.OrganizationRepEmailAddress, "organizationrepemailaddress", "", "Organization representative's contact email to send OV callback email to")

	return cmd, nil
}

func newCmdSSLResendApproverEmail() (*cobra.Command, error) {
	var opts api.SSLResendApproverEmailOptions
	var err error

	cmd := &cobra.Command{
		Use:   "resendapproveremail",
		Short: "Resends the approver email.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLResendApproverEmail(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "The unique certificate ID that you get after calling ssl.create command")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdSSLGetInfo() (*cobra.Command, error) {
	var opts api.SSLGetInfoOptions
	var err error

	cmd := &cobra.Command{
		Use:   "getinfo",
		Short: "Retrieves information about the requested SSL certificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLGetInfo(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "Unique ID of the SSL certificate")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Returncertificate, "returncertificate", "", "A flag for returning certificate in response")
	cmd.Flags().StringVar(&opts.Returntype, "returntype", "", "Type of returned certificate. Parameter takes 'Individual' (for X.509 format) or 'PKCS7' values")

	return cmd, nil
}

func newCmdSSLRenew() (*cobra.Command, error) {
	var opts api.SSLRenewOptions
	var err error

	cmd := &cobra.Command{
		Use:   "renew",
		Short: "Renews an SSL certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			yearsFlag := cmd.Flag("years").Value

			yearsVal, err := strconv.Atoi(yearsFlag.String())
			if err != nil {
				return err
			}

			if yearsVal != 1 && yearsVal != 2 {
				return errors.New("Year value must be 1 or 2.")
			}

			r, err := apiClient.SSLRenew(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "Unique ID of the SSL certificate you wish to renew")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.Years, "years", 0, "Number of years renewal SSL will be issued for. Allowed values: 1,2")
	if err = cmd.MarkFlagRequired("years"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.SSLType, "ssltype", "", "SSL product name.")
	if err = cmd.MarkFlagRequired("ssltype"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Promotional (coupon) code for the certificate")

	return cmd, nil
}

func newCmdSSLReissue() (*cobra.Command, error) {
	var opts api.SSLReissueOptions
	var HTTPDCValidation, DNSDCValidation bool
	var err error

	cmd := &cobra.Command{
		Use:   "reissue",
		Short: "Reissues an SSL certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpdcvFlag := cmd.Flag("httpdcvalidation")
			if httpdcvFlag.Changed {
				v, err := strconv.ParseBool(httpdcvFlag.Value.String())
				if err != nil {
					return err
				}
				opts.HTTPDCValidation = &v
			}

			r, err := apiClient.SSLReissue(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "Unique identifier of SSL certificate to be reissued")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.CSR, "csr", "", "Certificate Signing Request")
	if err = cmd.MarkFlagRequired("csr"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(
		&opts.AdminEmailAddress,
		"adminemailaddress",
		"",
		"Email address to send signed SSL certificate file to. It is not possible to change AdminEmailAddress in reissue due to limitation from Certificate Authority. SSL file will be sent to AdminEmailAddress used in the initial activate call.",
	)
	cmd.Flags().StringVar(
		&opts.WebServerType,
		"webservertype",
		"",
		"Server software where SSL will be installed. Defines SSL certificate file format (PEM or PKCS7).",
	)

	// Domain Control Validation (DCV)
	cmd.Flags().StringVar(&opts.ApproverEmail, "approveremail", "", "Sets method to pass DCV through.")
	cmd.Flags().BoolVar(&HTTPDCValidation, "httpdcvalidation", false, "Sets method to pass DCV through.")
	cmd.Flags().BoolVar(&DNSDCValidation, "dnsdcvalidation", false, "Sets method to pass DCV through.")

	// Multi-domain specific
	cmd.Flags().StringVar(&opts.DNSNames, "dnsnames", "", "Specifies domain names to be included into the certificate request in addition to the Common Name in (CSR) format")
	cmd.Flags().StringVar(&opts.DNSApproverEmails, "dnsapproveremails", "", "Sets method to pass DCV through for each domain in the certificate request")

	// OV & EV
	cmd.Flags().StringVar(&opts.AdminOrganizationName, "adminorganizationname", "", "Organization name to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.OrganizationDepartment, "organizationdepartment", "", "Organizational Unit Name")
	cmd.Flags().StringVar(&opts.AdminCountry, "admincountry", "", "Organization country name to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminStateProvince, "adminstateprovince", "", "Organization state or province to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminCity, "admincity", "", "Organization city or locality to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminAddress1, "adminaddress1", "", "Organization street address to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminAddress2, "adminaddress2", "", "Additional organization street address to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminPostalCode, "adminpostalcode", "", "Organization postal code to be validated by Comodo validation staff and encoded into the issued certificate")
	cmd.Flags().StringVar(&opts.AdminPhone, "adminphone", "", "Organization phone number registered in any reputable online database used for certificate request verification callback")
	cmd.Flags().StringVar(&opts.OrganizationDUNS, "organizationduns", "", "Organization's number in DUN & Bradstreet company database")

	// EV only
	cmd.Flags().StringVar(&opts.CompanyIncorporationCountry, "companyincorporationcountry", "", "Country where the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyIncorporationStateProvince, "companyincorporationstateprovince", "", "State or province where the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyIncorporationLocality, "companyincorporationlocality", "", "City or locality where the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyIncorporationDate, "companyincorporationdate", "", "Date when the organization’s legal existence was established")
	cmd.Flags().StringVar(&opts.CompanyDBA, "companydba", "", "Organization's 'Doing business as' name ")
	cmd.Flags().StringVar(&opts.CompanyRegistrationNumber, "companyregistrationnumber", "", "Number assigned to an organization by an Incorporating Agency")

	// OV only
	cmd.Flags().StringVar(&opts.OrganizationRepFirstName, "organizationrepfirstname", "", "Full-time organization representative's first name")
	cmd.Flags().StringVar(&opts.OrganizationRepLastName, "OrganizationRepLastName", "", "Full-time organization representative's last name")
	cmd.Flags().StringVar(&opts.OrganizationRepTitle, "organizationreptitle", "", "Full-time organization representative's job title")
	cmd.Flags().StringVar(&opts.OrganizationRepPhone, "organizationrepphone", "", "Full-time organization representative's phone number")
	cmd.Flags().StringVar(&opts.OrganizationRepEmailAddress, "organizationrepemailaddress", "", "Organization representative's contact email to send OV callback email to")

	return cmd, nil
}

func newCmdSSLResendFulfillmentEmail() (*cobra.Command, error) {
	var opts api.SSLResendFulfillmentEmailOptions
	var err error

	cmd := &cobra.Command{
		Use:   "resendfulfillmentemail",
		Short: "Resends the fulfillment email containing the certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLResendFulfillmentEmail(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "The unique certificate ID that you get after calling ssl.create command")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdSSLPurchaseMoreSans() (*cobra.Command, error) {
	var opts api.SSLPurchaseMoreSansOptions
	var err error

	cmd := &cobra.Command{
		Use:   "purchasemoresans",
		Short: "Purchases more add-on domains for an already purchased certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLPurchaseMoreSans(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.CertificateID, "certificateid", "", "ID of the certificate for which you wish to purchase more add-on domains")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.NumberOfSANSToAdd, "numberofsanstoadd", 0, "Number of add-on domains to be ordered")
	if err = cmd.MarkFlagRequired("numberofsanstoadd"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdSSLRevokeCertificate() (*cobra.Command, error) {
	var opts api.SSLRevokeCertificateOptions
	var err error

	cmd := &cobra.Command{
		Use:   "revokecertificate",
		Short: "Revokes a reissued SSL certificate.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLRevokeCertificate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.CertificateID, "certificateid", 0, "ID of the certificate for you wish to revoke")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.CertificateType, "certificatetype", "", "Type of SSL Certificate")
	if err = cmd.MarkFlagRequired("certificatetype"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdSSLEditDCVMethod() (*cobra.Command, error) {
	var opts api.SSLEditDCVMethodOptions
	var err error

	cmd := &cobra.Command{
		Use:   "editdcvmethod",
		Short: "Sets new domain control validation (DCV) method for a certificate or serves as 'retry' mechanism.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.SSLEditDCVMethod(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.CertificateID, "certificateid", 0, "Unique ID of the SSL certificate to set new DCV method for")
	if err = cmd.MarkFlagRequired("certificateid"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.DCVMethod, "dcvmethod", "", "DCV method to validate domain control with")
	cmd.Flags().StringVar(&opts.DNSNames, "dnsnames", "", "Comma-separated list of domain names to set new DCV method for")
	cmd.Flags().StringVar(&opts.DCVMethods, "dcvmethods", "", "Comma-separated list of DCV methods to validate domain control with")

	return cmd, nil
}
