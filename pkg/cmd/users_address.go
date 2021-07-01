package cmd

import (
	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdUsersAddress() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "Manage user's addresses.",
	}

	addCommand(cmd, newCmdUsersAddressCreate)
	addCommand(cmd, newCmdUsersAddressDelete)
	addCommand(cmd, newCmdUsersAddressGetInfo)
	addCommand(cmd, newCmdUsersAddressGetList)
	addCommand(cmd, newCmdUsersAddressSetDefault)

	return cmd, nil
}

func newCmdUsersAddressCreate() (*cobra.Command, error) {
	var opts api.UsersAddressCreateOptions
	var err error

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new address for the user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersAddressCreate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.AddressName, "addressname", "", "Address name to create.")
	if err = cmd.MarkFlagRequired("addressname"); err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.DefaultYN, "defaultyn", 0, "Possible values are 0 and 1. If the value of this parameter is set to 1, the address is set as default address for the user.")

	cmd.Flags().StringVar(&opts.EmailAddress, "emailaddress", "", "Email address of the user.")
	if err = cmd.MarkFlagRequired("emailaddress"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.FirstName, "firstname", "", "First name of the user.")
	if err = cmd.MarkFlagRequired("firstname"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.LastName, "lastname", "", "Last name of the user.")
	if err = cmd.MarkFlagRequired("lastname"); err != nil {
		return nil, err
	}

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

	cmd.Flags().StringVar(&opts.StateProvince, "stateprovince", "", "State/Province of the user.")
	if err = cmd.MarkFlagRequired("stateprovince"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.StateProvinceChoice, "stateprovincechoice", "", "State/Province choice of the user.")
	if err = cmd.MarkFlagRequired("stateprovincechoice"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Zip, "zip", "", "Zip/Postal code of the user.")
	if err = cmd.MarkFlagRequired("zip"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Country, "country", "", "Two letter country code of the user.")
	if err = cmd.MarkFlagRequired("country"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Phone, "phone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	if err = cmd.MarkFlagRequired("phone"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.PhoneExt, "phoneext", "", "PhoneExt of the user.")

	cmd.Flags().StringVar(&opts.Fax, "fax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	if err = cmd.MarkFlagRequired("fax"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersAddressDelete() (*cobra.Command, error) {
	var opts api.UsersAddressDeleteOptions
	var err error

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes the particular address for the user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersAddressDelete(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.AddressId, "addressid", 0, "The unique AddressID to delete.")
	if err = cmd.MarkFlagRequired("addressid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersAddressGetInfo() (*cobra.Command, error) {
	var opts api.UsersAddressGetInfoOptions
	var err error

	cmd := &cobra.Command{
		Use:   "getinfo",
		Short: "Gets information for the requested addressID.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersAddressGetInfo(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.AddressId, "addressid", 0, "The unique AddressID.")
	if err = cmd.MarkFlagRequired("addressid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersAddressGetList() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Gets a list of addressIDs and addressnames associated with the user account.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersAddressGetList()
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	return cmd, nil
}

func newCmdUsersAddressSetDefault() (*cobra.Command, error) {
	var opts api.UsersAddressSetDefaultOptions
	var err error

	cmd := &cobra.Command{
		Use:   "setdefault",
		Short: "Sets default address for the user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersAddressSetDefault(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.AddressId, "addressid", 0, "The unique addressID to set default.")
	if err = cmd.MarkFlagRequired("addressid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdUsersAddressUpdate() (*cobra.Command, error) {
	var opts api.UsersAddressUpdateOptions
	var err error

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the particular address of the user.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.UsersAddressUpdate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.AddressId, "addressid", 0, "The unique address ID to update.")
	if err = cmd.MarkFlagRequired("addressid"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.AddressName, "addressname", "", "Address name to create.")
	if err = cmd.MarkFlagRequired("addressname"); err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.DefaultYN, "defaultyn", 0, "Possible values are 0 and 1. If the value of this parameter is set to 1, the address is set as default address for the user.")

	cmd.Flags().StringVar(&opts.EmailAddress, "emailaddress", "", "Email address of the user.")
	if err = cmd.MarkFlagRequired("emailaddress"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.FirstName, "firstname", "", "First name of the user.")
	if err = cmd.MarkFlagRequired("firstname"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.LastName, "lastname", "", "Last name of the user.")
	if err = cmd.MarkFlagRequired("lastname"); err != nil {
		return nil, err
	}

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

	cmd.Flags().StringVar(&opts.StateProvince, "stateprovince", "", "State/Province of the user.")
	if err = cmd.MarkFlagRequired("stateprovince"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.StateProvinceChoice, "stateprovincechoice", "", "State/Province choice of the user.")
	if err = cmd.MarkFlagRequired("stateprovincechoice"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Zip, "zip", "", "Zip/Postal code of the user.")
	if err = cmd.MarkFlagRequired("zip"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Country, "country", "", "Two letter country code of the user.")
	if err = cmd.MarkFlagRequired("country"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Phone, "phone", "", "Phone number in the format +NNN.NNNNNNNNNN")
	if err = cmd.MarkFlagRequired("phone"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.PhoneExt, "phoneext", "", "PhoneExt of the user.")

	cmd.Flags().StringVar(&opts.Fax, "fax", "", "Fax number in the format +NNN.NNNNNNNNNN")
	if err = cmd.MarkFlagRequired("fax"); err != nil {
		return nil, err
	}

	return cmd, nil
}
