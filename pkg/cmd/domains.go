package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdDomains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domains",
		Short: "Manage domains",
	}

	addCommand(cmd, newCmdDomainsDNS)
	addCommand(cmd, newCmdDomainsNS)
	addCommand(cmd, newCmdDomainsGetList)
	addCommand(cmd, newCmdDomainsGetContacts)
	addCommand(cmd, newCmdDomainsCreate)
	addCommand(cmd, newCmdDomainsGetTldList)
	addCommand(cmd, newCmdDomainsSetContacts)
	addCommand(cmd, newCmdDomainsCheck)
	addCommand(cmd, newCmdDomainsReactivate)
	addCommand(cmd, newCmdDomainsRenew)
	addCommand(cmd, newCmdDomainsGetRegistrarLock)
	addCommand(cmd, newCmdDomainsSetRegistrarLock)
	addCommand(cmd, newCmdDomainsGetInfo)

	return cmd
}

func newCmdDomainsGetList() (*cobra.Command, error) {
	var opts api.DomainsGetListOptions

	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Returns a list of domains for the particular user",
		RunE: func(cmd *cobra.Command, arg []string) error {
			r, err := apiClient.DomainsGetList(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.ListType, "listtype", "ALL", "Possible values are ALL, EXPIRING, or EXPIRED")
	cmd.Flags().StringVar(&opts.SearchTerm, "searchterm", "", "Keyword to look for in the domain list")
	cmd.Flags().StringVar(&opts.Page, "page", "1", "Page to return")
	cmd.Flags().StringVar(&opts.PageSize, "pagesize", "20", "Number of domains to be listed on a page. Minimum value is 10, and maximum value is 100")
	cmd.Flags().StringVar(&opts.SortBy, "sortby", "", "Possible values are NAME, NAME_DESC, EXPIREDATE, EXPIREDATE_DESC, CREATEDATE, CREATEDATE_DESC")

	return cmd, nil
}

func newCmdDomainsGetContacts() (*cobra.Command, error) {
	var err error
	var opts api.DomainsGetContactsOptions

	cmd := &cobra.Command{
		Use:   "getcontacts",
		Short: "Gets contact information for the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsGetContacts(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain to get contacts")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsCreate() (*cobra.Command, error) {
	var err error
	var opts api.DomainsCreateOptions
	var reuseRegistrantInfo bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Registers a new domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			// add space to beginning of phonenumber so that URL encode
			// will add the correct + character that Namecheap requires
			opts.RegistrantPhone = fmt.Sprintf(" %s", opts.RegistrantPhone)

			if reuseRegistrantInfo {
				opts.TechFirstName = opts.RegistrantFirstName
				opts.TechLastName = opts.RegistrantLastName
				opts.TechAddress1 = opts.RegistrantAddress1
				opts.TechCity = opts.RegistrantCity
				opts.TechStateProvince = opts.RegistrantStateProvince
				opts.TechPostalCode = opts.RegistrantPostalCode
				opts.TechCountry = opts.RegistrantCountry
				opts.TechPhone = opts.RegistrantPhone
				opts.TechEmailAddress = opts.RegistrantEmailAddress

				opts.AdminFirstName = opts.RegistrantFirstName
				opts.AdminLastName = opts.RegistrantLastName
				opts.AdminAddress1 = opts.RegistrantAddress1
				opts.AdminCity = opts.RegistrantCity
				opts.AdminStateProvince = opts.RegistrantStateProvince
				opts.AdminPostalCode = opts.RegistrantPostalCode
				opts.AdminCountry = opts.RegistrantCountry
				opts.AdminPhone = opts.RegistrantPhone
				opts.AdminEmailAddress = opts.RegistrantEmailAddress

				opts.AuxBillingFirstName = opts.RegistrantFirstName
				opts.AuxBillingLastName = opts.RegistrantLastName
				opts.AuxBillingAddress1 = opts.RegistrantAddress1
				opts.AuxBillingCity = opts.RegistrantCity
				opts.AuxBillingStateProvince = opts.RegistrantStateProvince
				opts.AuxBillingPostalCode = opts.RegistrantPostalCode
				opts.AuxBillingCountry = opts.RegistrantCountry
				opts.AuxBillingPhone = opts.RegistrantPhone
				opts.AuxBillingEmailAddress = opts.RegistrantEmailAddress
			} else {
				opts.TechPhone = fmt.Sprintf(" %s", opts.TechPhone)
				opts.AdminPhone = fmt.Sprintf(" %s", opts.AdminPhone)
				opts.AuxBillingPhone = fmt.Sprintf(" %s", opts.AuxBillingPhone)
			}

			if opts.AddFreeWhoisguard != "" {
				if opts.AddFreeWhoisguard != "yes" && opts.AddFreeWhoisguard != "no" {
					cmd.SilenceUsage = true
					return errors.New("Invalid option for `addfreewhoisguard`. Accepted values: [yes, no].\n")
				}
			}
			if opts.WGEnabled != "" {
				if opts.WGEnabled != "yes" && opts.WGEnabled != "no" {
					cmd.SilenceUsage = true
					return errors.New("Invalid option for `wgenabled`. Accepted values: [yes, no].\n")
				}
			}

			r, err := apiClient.DomainsCreate(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().BoolVar(&reuseRegistrantInfo, "reuseregistrantinfo", false, "Reuse Registrant information for all other contact information [Tech, Admin, AuxBilling]")

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to register")
	cmd.Flags().StringVar(&opts.Years, "years", "", "Number of years to register")
	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Promotional (coupon) code for the domain")
	cmd.Flags().StringVar(&opts.IdnCode, "idncode", "", "Code of Internationalized Domain Name")
	cmd.Flags().StringVar(&opts.Nameservers, "nameservers", "", "Comma-separated list of custom nameservers to be associated with the domain name")
	cmd.Flags().StringVar(&opts.AddFreeWhoisguard, "addfreewhoisguard", "yes", "Adds free WhoisGuard for the domain")
	cmd.Flags().StringVar(&opts.WGEnabled, "wgenabled", "yes", "Enables free WhoisGuard for the domain ")
	cmd.Flags().BoolVar(&opts.IsPremiumDomain, "ispremiumdomain", false, "Indication if the domain name is premium")
	cmd.Flags().StringVar(&opts.PremiumPrice, "premiumprice", "", "Registration price for the premium domain")
	cmd.Flags().StringVar(&opts.EapFee, "eapfee", "", "Purchase fee for the premium domain during Early Access Program (EAP)")

	cmd.Flags().StringVar(&opts.RegistrantOrganizationName, "organizationname", "", "Organization of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantJobTitle, "jobtitle", "", "Job title of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantFirstName, "firstname", "", "First name of the Registrant user")
	err = cmd.MarkFlagRequired("firstname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantLastName, "lastname", "", "Second name of the Registrant user")
	err = cmd.MarkFlagRequired("lastname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantAddress1, "address1", "", "Address1 of the Registrant user")
	err = cmd.MarkFlagRequired("address1")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantAddress2, "address2", "", "Address2 of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantCity, "city", "", "City of the Registrant user")
	err = cmd.MarkFlagRequired("city")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantStateProvince, "stateprovince", "", "State/Province of the Registrant user")
	err = cmd.MarkFlagRequired("stateprovince")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantStateProvinceChoice, "stateprovincechoice", "", "StateProvinceChoice of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantPostalCode, "postalcode", "", "PostalCode of the Registrant user")
	err = cmd.MarkFlagRequired("postalcode")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantCountry, "country", "", "Country of the Registrant user")
	err = cmd.MarkFlagRequired("country")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantPhone, "phone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	err = cmd.MarkFlagRequired("phone")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantPhoneExt, "phoneext", "", "PhoneExt of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantFax, "fax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.RegistrantEmailAddress, "emailaddress", "", "Email address of the Registrant user")
	err = cmd.MarkFlagRequired("emailaddress")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.TechOrganizationName, "techorganizationname", "", "Organization of the Tech user")
	cmd.Flags().StringVar(&opts.TechJobTitle, "techjobtitle", "", "Job title of the Tech user")
	cmd.Flags().StringVar(&opts.TechFirstName, "techfirstname", "", "First name of the Tech user")
	cmd.Flags().StringVar(&opts.TechLastName, "techlastname", "", "Second name of the Tech user")
	cmd.Flags().StringVar(&opts.TechAddress1, "techaddress1", "", "Address1 of the Tech user")
	cmd.Flags().StringVar(&opts.TechAddress2, "techaddress2", "", "Address2 of the Tech user")
	cmd.Flags().StringVar(&opts.TechCity, "techcity", "", "City of the Tech user")
	cmd.Flags().StringVar(&opts.TechStateProvince, "techstateprovince", "", "State/Province of the Tech user")
	cmd.Flags().StringVar(&opts.TechStateProvinceChoice, "techstateprovincechoice", "", "StateProvinceChoice of the Tech user")
	cmd.Flags().StringVar(&opts.TechPostalCode, "techpostalcode", "", "PostalCode of the Tech user")
	cmd.Flags().StringVar(&opts.TechCountry, "techcountry", "", "Country of the Tech user")
	cmd.Flags().StringVar(&opts.TechPhone, "techphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.TechPhoneExt, "techphoneext", "", "PhoneExt of the Tech user")
	cmd.Flags().StringVar(&opts.TechFax, "techfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.TechEmailAddress, "techemailaddress", "", "Email address of the Tech user")

	cmd.Flags().StringVar(&opts.AdminOrganizationName, "adminorganizationname", "", "Organization of the Admin user")
	cmd.Flags().StringVar(&opts.AdminJobTitle, "adminjobtitle", "", "Job title of the Admin user")
	cmd.Flags().StringVar(&opts.AdminFirstName, "adminfirstname", "", "First name of the Admin user")
	cmd.Flags().StringVar(&opts.AdminLastName, "adminlastname", "", "Second name of the Admin user")
	cmd.Flags().StringVar(&opts.AdminAddress1, "adminaddress1", "", "Address1 of the Admin user")
	cmd.Flags().StringVar(&opts.AdminAddress2, "adminaddress2", "", "Address2 of the Admin user")
	cmd.Flags().StringVar(&opts.AdminCity, "admincity", "", "City of the Admin user")
	cmd.Flags().StringVar(&opts.AdminStateProvince, "adminstateprovince", "", "State/Province of the Admin user")
	cmd.Flags().StringVar(&opts.AdminStateProvinceChoice, "adminstateprovincechoice", "", "StateProvinceChoice of the Admin user")
	cmd.Flags().StringVar(&opts.AdminPostalCode, "adminpostalcode", "", "PostalCode of the Admin user")
	cmd.Flags().StringVar(&opts.AdminCountry, "admincountry", "", "Country of the Admin user")
	cmd.Flags().StringVar(&opts.AdminPhone, "adminphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AdminPhoneExt, "adminphoneext", "", "PhoneExt of the Admin user")
	cmd.Flags().StringVar(&opts.AdminFax, "adminfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AdminEmailAddress, "adminemailaddress", "", "Email address of the Admin user")

	cmd.Flags().StringVar(&opts.AuxBillingOrganizationName, "auxbillingorganizationname", "", "Organization of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingJobTitle, "auxbillingjobtitle", "", "Job title of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingFirstName, "auxbillingfirstname", "", "First name of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingLastName, "auxbillinglastname", "", "Second name of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingAddress1, "auxbillingaddress1", "", "Address1 of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingAddress2, "auxbillingaddress2", "", "Address2 of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingCity, "auxbillingcity", "", "City of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingStateProvince, "auxbillingstateprovince", "", "State/Province of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingStateProvinceChoice, "auxbillingstateprovincechoice", "", "StateProvinceChoice of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingPostalCode, "auxbillingpostalcode", "", "PostalCode of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingCountry, "auxbillingcountry", "", "Country of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingPhone, "auxbillingphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AuxBillingPhoneExt, "auxbillingphoneext", "", "PhoneExt of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingFax, "auxbillingfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AuxBillingEmailAddress, "auxbillingemailaddress", "", "Email address of the AuxBilling user")

	cmd.Flags().StringVar(&opts.BillingFirstName, "billingfirstname", "", "First name of the Billing user")
	cmd.Flags().StringVar(&opts.BillingLastName, "billinglastname", "", "Last name of the Billing user")
	cmd.Flags().StringVar(&opts.BillingAddress1, "billingaddress1", "", "Address1 of the Billing user")
	cmd.Flags().StringVar(&opts.BillingAddress2, "billingaddress2", "", "Address2 of the Billing user")
	cmd.Flags().StringVar(&opts.BillingCity, "billingcity", "", "City of the Billing user")
	cmd.Flags().StringVar(&opts.BillingStateProvince, "billingstateprovince", "", "State/Province of the Billing user")
	cmd.Flags().StringVar(&opts.BillingStateProvinceChoice, "billingstateprovincechoice", "", "StateProvinceChoice of the Billing user")
	cmd.Flags().StringVar(&opts.BillingPostalCode, "billingpostalcode", "", "PostalCode of the Billing user")
	cmd.Flags().StringVar(&opts.BillingCountry, "billingcountry", "", "Country of the Billing user")
	cmd.Flags().StringVar(&opts.BillingPhone, "billingphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.BillingPhoneExt, "billingphoneext", "", "PhoneExt of the Billing user")
	cmd.Flags().StringVar(&opts.BillingFax, "billingfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.BillingEmailAddress, "billingemailaddress", "", "Email address of the Billing user")

	// Extended attributes ?

	return cmd, nil
}

func newCmdDomainsGetTldList() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "gettldlist",
		Short: "Returns a list of TLDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsGetTldList()
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	return cmd, nil
}

func newCmdDomainsSetContacts() (*cobra.Command, error) {
	var err error
	var opts api.DomainsSetContactsOptions
	var reuseRegistrantInfo bool

	cmd := &cobra.Command{
		Use:   "setcontacts",
		Short: "Sets contact information for the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if reuseRegistrantInfo {
				opts.TechFirstName = opts.RegistrantFirstName
				opts.TechLastName = opts.RegistrantLastName
				opts.TechAddress1 = opts.RegistrantAddress1
				opts.TechCity = opts.RegistrantCity
				opts.TechStateProvince = opts.RegistrantStateProvince
				opts.TechPostalCode = opts.RegistrantPostalCode
				opts.TechCountry = opts.RegistrantCountry
				opts.TechPhone = opts.RegistrantPhone
				opts.TechEmailAddress = opts.RegistrantEmailAddress

				opts.AdminFirstName = opts.RegistrantFirstName
				opts.AdminLastName = opts.RegistrantLastName
				opts.AdminAddress1 = opts.RegistrantAddress1
				opts.AdminCity = opts.RegistrantCity
				opts.AdminStateProvince = opts.RegistrantStateProvince
				opts.AdminPostalCode = opts.RegistrantPostalCode
				opts.AdminCountry = opts.RegistrantCountry
				opts.AdminPhone = opts.RegistrantPhone
				opts.AdminEmailAddress = opts.RegistrantEmailAddress

				opts.AuxBillingFirstName = opts.RegistrantFirstName
				opts.AuxBillingLastName = opts.RegistrantLastName
				opts.AuxBillingAddress1 = opts.RegistrantAddress1
				opts.AuxBillingCity = opts.RegistrantCity
				opts.AuxBillingStateProvince = opts.RegistrantStateProvince
				opts.AuxBillingPostalCode = opts.RegistrantPostalCode
				opts.AuxBillingCountry = opts.RegistrantCountry
				opts.AuxBillingPhone = opts.RegistrantPhone
				opts.AuxBillingEmailAddress = opts.RegistrantEmailAddress

			}

			r, err := apiClient.DomainsSetContacts(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().BoolVar(&reuseRegistrantInfo, "reuseregistrantinfo", false, "Reuse Registrant information for all other contact information [Tech, Admin, AuxBilling]")

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "The domain name to register")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantOrganizationName, "organizationname", "", "Organization of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantJobTitle, "jobtitle", "", "Job title of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantFirstName, "firstname", "", "First name of the Registrant user")
	err = cmd.MarkFlagRequired("firstname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantLastName, "lastname", "", "Second name of the Registrant user")
	err = cmd.MarkFlagRequired("lastname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantAddress1, "address1", "", "Address1 of the Registrant user")
	err = cmd.MarkFlagRequired("address1")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantAddress2, "address2", "", "Address2 of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantCity, "city", "", "City of the Registrant user")
	err = cmd.MarkFlagRequired("city")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantStateProvince, "stateprovince", "", "State/Province of the Registrant user")
	err = cmd.MarkFlagRequired("stateprovince")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantStateProvinceChoice, "stateprovincechoice", "", "StateProvinceChoice of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantPostalCode, "postalcode", "", "PostalCode of the Registrant user")
	err = cmd.MarkFlagRequired("postalcode")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantCountry, "country", "", "Country of the Registrant user")
	err = cmd.MarkFlagRequired("country")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantPhone, "phone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	err = cmd.MarkFlagRequired("phone")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.RegistrantPhoneExt, "phoneext", "", "PhoneExt of the Registrant user")
	cmd.Flags().StringVar(&opts.RegistrantFax, "fax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.RegistrantEmailAddress, "emailaddress", "", "Email address of the Registrant user")
	err = cmd.MarkFlagRequired("emailaddress")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.TechOrganizationName, "techorganizationname", "", "Organization of the Tech user")
	cmd.Flags().StringVar(&opts.TechJobTitle, "techjobtitle", "", "Job title of the Tech user")
	cmd.Flags().StringVar(&opts.TechFirstName, "techfirstname", "", "First name of the Tech user")
	cmd.Flags().StringVar(&opts.TechLastName, "techlastname", "", "Second name of the Tech user")
	cmd.Flags().StringVar(&opts.TechAddress1, "techaddress1", "", "Address1 of the Tech user")
	cmd.Flags().StringVar(&opts.TechAddress2, "techaddress2", "", "Address2 of the Tech user")
	cmd.Flags().StringVar(&opts.TechCity, "techcity", "", "City of the Tech user")
	cmd.Flags().StringVar(&opts.TechStateProvince, "techstateprovince", "", "State/Province of the Tech user")
	cmd.Flags().StringVar(&opts.TechStateProvinceChoice, "techstateprovincechoice", "", "StateProvinceChoice of the Tech user")
	cmd.Flags().StringVar(&opts.TechPostalCode, "techpostalcode", "", "PostalCode of the Tech user")
	cmd.Flags().StringVar(&opts.TechCountry, "techcountry", "", "Country of the Tech user")
	cmd.Flags().StringVar(&opts.TechPhone, "techphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.TechPhoneExt, "techphoneext", "", "PhoneExt of the Tech user")
	cmd.Flags().StringVar(&opts.TechFax, "techfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.TechEmailAddress, "techemailaddress", "", "Email address of the Tech user")

	cmd.Flags().StringVar(&opts.AdminOrganizationName, "adminorganizationname", "", "Organization of the Admin user")
	cmd.Flags().StringVar(&opts.AdminJobTitle, "adminjobtitle", "", "Job title of the Admin user")
	cmd.Flags().StringVar(&opts.AdminFirstName, "adminfirstname", "", "First name of the Admin user")
	cmd.Flags().StringVar(&opts.AdminLastName, "adminlastname", "", "Second name of the Admin user")
	cmd.Flags().StringVar(&opts.AdminAddress1, "adminaddress1", "", "Address1 of the Admin user")
	cmd.Flags().StringVar(&opts.AdminAddress2, "adminaddress2", "", "Address2 of the Admin user")
	cmd.Flags().StringVar(&opts.AdminCity, "admincity", "", "City of the Admin user")
	cmd.Flags().StringVar(&opts.AdminStateProvince, "adminstateprovince", "", "State/Province of the Admin user")
	cmd.Flags().StringVar(&opts.AdminStateProvinceChoice, "adminstateprovincechoice", "", "StateProvinceChoice of the Admin user")
	cmd.Flags().StringVar(&opts.AdminPostalCode, "adminpostalcode", "", "PostalCode of the Admin user")
	cmd.Flags().StringVar(&opts.AdminCountry, "admincountry", "", "Country of the Admin user")
	cmd.Flags().StringVar(&opts.AdminPhone, "adminphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AdminPhoneExt, "adminphoneext", "", "PhoneExt of the Admin user")
	cmd.Flags().StringVar(&opts.AdminFax, "adminfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AdminEmailAddress, "adminemailaddress", "", "Email address of the Admin user")

	cmd.Flags().StringVar(&opts.AuxBillingOrganizationName, "auxbillingorganizationname", "", "Organization of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingJobTitle, "auxbillingjobtitle", "", "Job title of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingFirstName, "auxbillingfirstname", "", "First name of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingLastName, "auxbillinglastname", "", "Second name of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingAddress1, "auxbillingaddress1", "", "Address1 of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingAddress2, "auxbillingaddress2", "", "Address2 of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingCity, "auxbillingcity", "", "City of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingStateProvince, "auxbillingstateprovince", "", "State/Province of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingStateProvinceChoice, "auxbillingstateprovincechoice", "", "StateProvinceChoice of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingPostalCode, "auxbillingpostalcode", "", "PostalCode of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingCountry, "auxbillingcountry", "", "Country of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingPhone, "auxbillingphone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AuxBillingPhoneExt, "auxbillingphoneext", "", "PhoneExt of the AuxBilling user")
	cmd.Flags().StringVar(&opts.AuxBillingFax, "auxbillingfax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	cmd.Flags().StringVar(&opts.AuxBillingEmailAddress, "auxbillingemailaddress", "", "Email address of the AuxBilling user")

	// Extended attributes ?

	return cmd, nil
}

func newCmdDomainsCheck() (*cobra.Command, error) {
	var err error
	var opts api.DomainsCheckOptions

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Checks the availability of domains",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsCheck(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainList, "domainlist", "", "Comma-separated list of domains to check")
	err = cmd.MarkFlagRequired("domainlist")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsReactivate() (*cobra.Command, error) {
	var opts api.DomainsReactivateOptions
	var isPremiumDomain, yearsToAdd string

	cmd := &cobra.Command{
		Use:   "reactivate",
		Short: "Reactivates an expired domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if isPremiumDomain != "" {
				value, err := strconv.ParseBool(isPremiumDomain)
				if err != nil {
					return errors.New("Invalid option for `ispremiumdomain`. Accepted values: [true, false].\n")
				}

				opts.IsPremiumDomain = value
			}

			if yearsToAdd != "" {
				value, err := strconv.Atoi(yearsToAdd)
				if err != nil {
					return errors.New("Invalid option for `yearstoadd`. Must be valid number.\n")
				}

				opts.YearsToAdd = value
			}

			r, err := apiClient.DomainsReactivate(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to reactivate")
	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Promotional (coupon) code for reactivating the domain")
	cmd.Flags().StringVar(&opts.PremiumPrice, "premiumprice", "", "Reactivation price for the premium domain")
	cmd.Flags().StringVar(&yearsToAdd, "yearstoadd", "", "Number of years after expiry")
	cmd.Flags().StringVar(&isPremiumDomain, "ispremiumdomain", "", "Indication if the domain name is premium")

	return cmd, nil
}

func newCmdDomainsRenew() (*cobra.Command, error) {
	var err error
	var opts api.DomainRenewOptions
	var isPremiumDomain string

	cmd := &cobra.Command{
		Use:   "renew",
		Short: "Renews an expiring domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if isPremiumDomain != "" {
				value, err := strconv.ParseBool(isPremiumDomain)
				if err != nil {
					return errors.New("Invalid option for `ispremiumdomain`. Accepted values: [true, false].\n")
				}

				opts.IsPremiumDomain = value
			}

			if opts.Years < 1 {
				return errors.New("Invalid option for `years`. Must be an integer greater than 0.\n")
			}

			r, err := apiClient.DomainsRenew(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to renew")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().IntVar(&opts.Years, "years", 0, "Number of years to renew")
	err = cmd.MarkFlagRequired("years")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Domain name to renew")
	cmd.Flags().StringVar(&opts.PremiumPrice, "premiumprice", "", "Renewal price for the premium domain")
	cmd.Flags().StringVar(&isPremiumDomain, "ispremiumdomain", "", "Indication if the domain name is premium")

	return cmd, nil
}

func newCmdDomainsGetRegistrarLock() (*cobra.Command, error) {
	var err error
	var opts api.DomainsGetRegistrarLockOptions

	cmd := &cobra.Command{
		Use:   "getregistrarlock",
		Short: "Gets the Registrar Lock status for the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsGetRegistrarLock(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to get status for")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsSetRegistrarLock() (*cobra.Command, error) {
	var err error
	var opts api.DomainsSetRegistrarLockOptions

	cmd := &cobra.Command{
		Use:   "setregistrarlock",
		Short: "Sets the Registrar Lock status for a domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.LockAction != "" {
				if opts.LockAction != "LOCK" && opts.LockAction != "UNLOCK" {
					return errors.New("Invalid option for `lockaction`. Accepted values: [LOCK, UNLOCK].")
				}
			}

			r, err := apiClient.DomainsSetRegistrarLock(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to set status for")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}
	cmd.Flags().StringVar(&opts.LockAction, "lockaction", "", "Possible values: LOCK, UNLOCK")

	return cmd, nil
}

func newCmdDomainsGetInfo() (*cobra.Command, error) {
	var err error
	var opts api.DomainsGetInfoOptions

	cmd := &cobra.Command{
		Use:   "getinfo",
		Short: "Returns information about the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsGetInfo(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name for which domain information needs to be requested")
	cmd.Flags().StringVar(&opts.HostName, "hostname", "", "Hosted domain name for which domain information needs to be requested")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
