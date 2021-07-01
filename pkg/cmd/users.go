package cmd

import (
	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdUsers() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Manage users",
	}

	addCommand(cmd, newCmdUsersAddress)
	addCommand(cmd, newCmdUsersGetPricing)
	addCommand(cmd, newCmdUsersGetBalances)
	addCommand(cmd, newCmdUsersChangePassword)
	addCommand(cmd, newCmdUsersUpdate)
	addCommand(cmd, newCmdUsersCreateAddFundsRequest)
	addCommand(cmd, newCmdUsersGetAddFundsStatus)
	addCommand(cmd, newCmdUsersCreate)
	addCommand(cmd, newCmdUsersLogin)
	addCommand(cmd, newCmdUsersResetPassword)

	return cmd, nil
}

func newCmdUsersGetPricing() (*cobra.Command, error) {
	var opts api.UsersGetPricingOptions
	var err error

	cmd := &cobra.Command{
		Use:   "getpricing",
		Short: "Returns pricing information for a requested product type.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersGetPricing(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.ProductType, "producttype", "", "Product Type to get pricing information")
	if err = cmd.MarkFlagRequired("producttype"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.ProductCategory, "productcategory", "", "Specific category within a product type")
	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Promotional (coupon) code for the user")
	cmd.Flags().StringVar(&opts.ActionName, "actionname", "", "Specific action within a product type")
	cmd.Flags().StringVar(&opts.ProductName, "productname", "", "The name of the product within a product type")

	return cmd, nil
}

func newCmdUsersGetBalances() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getbalances",
		Short: "Gets information about fund in the user's account.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersGetBalances()
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	return cmd, nil
}

func newCmdUsersChangePassword() (*cobra.Command, error) {
	var opts api.UsersChangePasswordOptions
	var err error

	cmd := &cobra.Command{
		Use:   "changepassword",
		Short: "Change user password.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersChangePassword(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.OldPassword, "oldpassword", "", "Old password of the user")
	if err = cmd.MarkFlagRequired("oldpassword"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.NewPassword, "newpassword", "", "New password of the user")
	if err = cmd.MarkFlagRequired("newpassword"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersUpdate() (*cobra.Command, error) {
	var opts api.UsersUpdateOptions
	var err error

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates user account information for the particular user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersUpdate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.FirstName, "firstname", "", "First name of the user")
	if err = cmd.MarkFlagRequired("firstname"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.LastName, "lastname", "", "Last name of the user")
	if err = cmd.MarkFlagRequired("lastname"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.JobTitle, "jobtitle", "", "Job designation of the user")
	cmd.Flags().StringVar(&opts.Organization, "organization", "", "Organization of the user")

	cmd.Flags().StringVar(&opts.Address1, "address1", "", "StreetAddress1 of the user")
	if err = cmd.MarkFlagRequired("address1"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Address2, "address2", "", "StreetAddress2 of the user")

	cmd.Flags().StringVar(&opts.City, "city", "", "City of the user")
	if err = cmd.MarkFlagRequired("city"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.StateProvince, "stateprovince", "", "State/Province of the user")
	if err = cmd.MarkFlagRequired("stateprovince"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Zip, "zip", "", "Zip/Postal code of the user")
	if err = cmd.MarkFlagRequired("zip"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Country, "country", "", "Two letter country code of the user")
	if err = cmd.MarkFlagRequired("country"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.EmailAddress, "emailaddress", "", "Email address of the user")
	if err = cmd.MarkFlagRequired("emailaddress"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Phone, "phone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	if err = cmd.MarkFlagRequired("phone"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.PhoneExt, "phoneext", "", "PhoneExt of the user")
	cmd.Flags().StringVar(&opts.Fax, "fax", "", "Fax number in the format +NNN.NNNNNNNNNN")

	return cmd, nil
}

func newCmdUsersCreateAddFundsRequest() (*cobra.Command, error) {
	var opts api.UsersCreateAddFundsRequestOptions
	var err error

	cmd := &cobra.Command{
		Use:   "createaddfundsrequest",
		Short: "Creates a request to add funds through a credit card.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersCreateAddFundsRequest(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.Username, "username", "", "Username to add funds to")
	if err = cmd.MarkFlagRequired("username"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.PaymentType, "paymenttype", "", "Allowed payment value: Creditcard")
	if err = cmd.MarkFlagRequired("paymenttype"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Amount, "amount", "", "Amount to add")
	if err = cmd.MarkFlagRequired("amount"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.ReturnUrl, "returnurl", "", "A valid URL to which the user should be redirected once payment is complete")
	if err = cmd.MarkFlagRequired("returnurl"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersGetAddFundsStatus() (*cobra.Command, error) {
	var opts api.UsersGetAddFundsStatusOptions
	var err error

	cmd := &cobra.Command{
		Use:   "getaddfundsstatus",
		Short: "Gets the status of add funds request.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersGetAddFundsStatus(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.TokenId, "tokenid", "", "The Unique ID that you received after calling namecheap.users.createaddfundsrequest method")
	if err = cmd.MarkFlagRequired("tokenid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersCreate() (*cobra.Command, error) {
	var opts api.UsersCreateOptions
	var err error

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new account at NameCheap under this ApiUser.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersCreate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.NewUserName, "newusername", "", "Username for the new user account")
	if err = cmd.MarkFlagRequired("newusername"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.NewUserPassword, "newuserpassword", "", "Password for the new user account")
	if err = cmd.MarkFlagRequired("newuserpassword"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.EmailAddress, "emailaddress", "", "Email address of the user")
	if err = cmd.MarkFlagRequired("emailaddress"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.FirstName, "firstname", "", "First name of the user")
	if err = cmd.MarkFlagRequired("firstname"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.LastName, "lastname", "", "Last name of the user")
	if err = cmd.MarkFlagRequired("lastname"); err != nil {
		return nil, err
	}

	cmd.Flags().BoolVar(&opts.IgnoreDuplicateEmailAddress, "ignoreduplicateemailaddress", true, "By default, it ignores to check if the email address is already in use.")
	cmd.Flags().IntVar(&opts.AcceptTerms, "acceptterms", 1, "Terms of agreement. The Value should be 1 for user account creation.")
	cmd.Flags().IntVar(&opts.AcceptNews, "acceptnews", 0, "Possible values are 0 and 1. Value should be 1 if the user wants to recieve newsletters else it should be 0.")
	cmd.Flags().StringVar(&opts.JobTitle, "jobtitle", "", "Job designation of the user.")
	cmd.Flags().StringVar(&opts.Organization, "organization", "", "Organization of the user.")

	cmd.Flags().StringVar(&opts.Address1, "address1", "", "StreetAddress1 of the user.")
	if err = cmd.MarkFlagRequired("address1"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Address2, "address2", "", "StreetAddress2 of the user.")

	cmd.Flags().StringVar(&opts.City, "city", "", "City of the user.")
	if err = cmd.MarkFlagRequired("city"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.StateProvince, "stateprovince", "", "StateProvince of the user.")
	if err = cmd.MarkFlagRequired("stateprovince"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Zip, "zip", "", "Zip of the user.")
	if err = cmd.MarkFlagRequired("zip"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Country, "country", "", "Country of the user.")
	if err = cmd.MarkFlagRequired("country"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Phone, "phone", "", "Phone number in the format +NNN.NNNNNNNNNN.")
	if err = cmd.MarkFlagRequired("phone"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.PhoneExt, "phoneext", "", "PhoneExt of the user.")
	cmd.Flags().StringVar(&opts.Fax, "fax", "", "Fax number in the format +NNN.NNNNNNNNNN.")

	return cmd, nil
}

func newCmdUsersLogin() (*cobra.Command, error) {
	var opts api.UsersLoginOptions
	var err error

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Validates the username and password of user accounts you have created using the API command namecheap.users.create.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersLogin(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.Password, "password", "", "Password of the user account.")
	if err = cmd.MarkFlagRequired("password"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersResetPassword() (*cobra.Command, error) {
	var opts api.UsersResetPasswordOptions
	var err error

	cmd := &cobra.Command{
		Use:   "resetpassword",
		Short: "When you call this API, a link to reset password will be emailed to the end user's profile email id. The end user needs to click on the link to reset password.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersResetPassword(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.FindBy, "findby", "", "Possible values: EMAILADDRESS, DOMAINNAME, USERNAME")
	if err = cmd.MarkFlagRequired("findby"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.FindByValue, "findbyvalue", "", "The username/email address/domain of the user.")
	if err = cmd.MarkFlagRequired("findbyvalue"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.EmailFromName, "emailfromname", "", "Enter a different value to overwrite the default value. Default Value: namecheap.com")
	cmd.Flags().StringVar(&opts.EmailFrom, "emailfrom", "", "Enter a different value to overwrite the default value. Default Value: support@namecheap.com ")
	cmd.Flags().StringVar(&opts.URLPattern, "urlpattern", "", "Enter a different URL to overwrite namecheap.com. Default Value: http://namecheap.com [RESETCODE] ")

	return cmd, nil
}
