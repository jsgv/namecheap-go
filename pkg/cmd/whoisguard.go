package cmd

import (
	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdWhoisguard() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "whoisguard",
		Short: "Manage whoisguard",
	}

	addCommand(cmd, newCmdWhoisguardChangeEmailAddress)
	addCommand(cmd, newCmdWhoisguardEnable)
	addCommand(cmd, newCmdWhoisguardDisable)
	addCommand(cmd, newCmdWhoisguardGetList)
	addCommand(cmd, newCmdWhoisguardRenew)

	return cmd, nil
}

func newCmdWhoisguardChangeEmailAddress() (*cobra.Command, error) {
	var opts api.WhoisguardChangeEmailAddressOptions
	var err error

	cmd := &cobra.Command{
		Use:   "changeemailaddress",
		Short: "Changes domain privacy email address.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.WhoisguardChangeEmailAddress(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.WhoisguardID, "whoisguardid", 0, "The unique domain privacy ID that you wish to change.")
	if err = cmd.MarkFlagRequired("whoisguardid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdWhoisguardEnable() (*cobra.Command, error) {
	var opts api.WhoisguardEnableOptions
	var err error

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enables domain privacy protection for the WhoisguardID.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.WhoisguardEnable(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.WhoisguardID, "whoisguardid", 0, "The unique domain privacy ID which you want to enable.")
	if err = cmd.MarkFlagRequired("whoisguardid"); err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.ForwardedToEmail, "forwardedtoemail", "", "The email address to which domain privacy emails are to be forwarded.")
	if err = cmd.MarkFlagRequired("forwardedtoemail"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdWhoisguardDisable() (*cobra.Command, error) {
	var opts api.WhoisguardDisableOptions
	var err error

	cmd := &cobra.Command{
		Use:   "disable",
		Short: "Disables domain privacy protection for the WhoisguardID.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.WhoisguardDisable(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().IntVar(&opts.WhoisguardID, "whoisguardid", 0, "The unique domain privacy ID which has to be disabled.")
	if err = cmd.MarkFlagRequired("whoisguardid"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdWhoisguardGetList() (*cobra.Command, error) {
	var opts api.WhoisguardGetListOptions

	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Gets the list of domain privacy protection.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.WhoisguardGetList(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.ListType, "listtype", "ALL", "Possible values are ALL, ALLOTED, FREE, DISCARD.")
	cmd.Flags().IntVar(&opts.Page, "page", 1, "Page to return.")
	cmd.Flags().IntVar(&opts.PageSize, "pagesize", 20, "Number of domain privacy subscriptions to be listed on a page. Minimum value: 2; Maximum value: 100.")

	return cmd, nil
}

func newCmdWhoisguardRenew() (*cobra.Command, error) {
	var opts api.WhoisguardRenewOptions
	var err error

	cmd := &cobra.Command{
		Use:   "renew",
		Short: "Renews domain privacy protection.",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.WhoisguardRenew(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.WhoisguardID, "whoisguardid", "", "Domain privacy ID to renew.")
	if err = cmd.MarkFlagRequired("whoisguardid"); err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.Years, "years", 1, "Number of years to renew")
	if err = cmd.MarkFlagRequired("years"); err != nil {
		return nil, err
	}

	cmd.Flags().IntVar(&opts.PromotionCode, "promotioncode", 1, "Promotional (coupon) code for renewing the domain privacy")

	return cmd, nil
}
