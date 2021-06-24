package cmd

import (
	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdDomainsTransfer() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "transfer",
		Short: "Manage domains transfers",
	}

	addCommand(cmd, newCmdDomainsTransferCreate)
	addCommand(cmd, newCmdDomainsTransferGetStatus)
	addCommand(cmd, newCmdDomainsTransferUpdateStatus)
	addCommand(cmd, newCmdDomainsTransferGetList)

	return cmd, nil
}

func newCmdDomainsTransferCreate() (*cobra.Command, error) {
	var err error
	var opts api.DomainsTransferCreateOptions

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Transfers a domain to Namecheap",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsTransferCreate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name to transfer")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Years, "years", "", "Years")
	err = cmd.MarkFlagRequired("years")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.EPPCode, "eppcode", "", "EPPCode is required for transferring certain domains")
	cmd.Flags().StringVar(&opts.PromotionCode, "promotioncode", "", "Promotional (coupon) code for transfer")
	cmd.Flags().BoolVar(&opts.AddFreeWhoisguard, "addfreewhoisguard", true, "Adds free Whoisguard for the domain")
	cmd.Flags().BoolVar(&opts.WGenable, "wgenable", true, "Enables free WhoisGuard for the domain")

	return cmd, nil
}

func newCmdDomainsTransferGetStatus() (*cobra.Command, error) {
	var err error
	var opts api.DomainsTransferGetStatusOptions

	cmd := &cobra.Command{
		Use:   "getstatus",
		Short: "Gets the status of a particular transfer.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsTransferGetStatus(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.TransferID, "transferid", 0, "The unique Transfer ID which you get after placing a transfer request")
	err = cmd.MarkFlagRequired("transferid")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsTransferUpdateStatus() (*cobra.Command, error) {
	var err error
	var opts api.DomainsTransferUpdateStatusOptions

	cmd := &cobra.Command{
		Use:   "updatestatus",
		Short: "Updates the status of a particular transfer. Allows you to re-submit the transfer after releasing the registry lock.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsTransferUpdateStatus(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.TransferID, "transferid", 0, "The unique Transfer ID which you get after placing a transfer request")
	err = cmd.MarkFlagRequired("transferid")
	if err != nil {
		return nil, err
	}

	cmd.Flags().BoolVar(&opts.Resubmit, "resubmit", true, "The value 'true' resubmits the transfer")

	return cmd, nil
}

func newCmdDomainsTransferGetList() (*cobra.Command, error) {
	var opts api.DomainsTransferGetListOptions

	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Gets the list of domain transfers.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsTransferGetList(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.ListType, "listtype", "", "Possible values are ALL, INPROGRESS, CANCELLED, COMPLETED")
	cmd.Flags().StringVar(&opts.SearchTerm, "searchterm", "", "The keyword should be a domain name")
	cmd.Flags().IntVar(&opts.Page, "page", 1, "Page to return")
	cmd.Flags().IntVar(&opts.PageSize, "pagesize", 10, "Number of transfer to be listed on a page")
	cmd.Flags().StringVar(&opts.SortBy, "sortby", "", "Possible values are DOMAINNAME, DOMAINNAME_DESC, TRANSFERDATE, TRANSFERDATE_DESC, STATUSDATE, STATUSDATE_DESC")

	return cmd, nil
}
