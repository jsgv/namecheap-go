package cmd

import (
	"errors"
	"strings"

	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdDomainsDNS() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Manage domains dns",
	}

	addCommand(cmd, newCmdDomainsDNSSetDefault)
	addCommand(cmd, newCmdDomainsDNSSetCustom)
	addCommand(cmd, newCmdDomainsDNSGetList)
	addCommand(cmd, newCmdDomainsDNSGetHosts)
	addCommand(cmd, newCmdDomainsDNSGetEmailForwarding)
	addCommand(cmd, newCmdDomainsDNSSetEmailForwarding)
	addCommand(cmd, newCmdDomainsDNSSetHosts)

	return cmd, nil
}

func newCmdDomainsDNSSetDefault() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSSetDefaultOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "setdefault",
		Short: "Sets domain to use Namecheap default DNS servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsDNSSetDefault(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsDNSSetCustom() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSSetCustomOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "setcustom",
		Short: "Sets domain to use custom DNS servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsDNSSetCustom(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Nameservers, "nameservers", "", "A comma-separated list of name servers to be associated with this domain")
	err = cmd.MarkFlagRequired("nameservers")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsDNSGetList() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSGetListOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Gets a list of DNS servers associated with the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsDNSGetList(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsDNSGetHosts() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSGetHostsOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "gethosts",
		Short: "Retrieves DNS host record settings for the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsDNSGetHosts(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsDNSGetEmailForwarding() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSGetEmailForwardingOptions

	cmd := &cobra.Command{
		Use:   "getemailforwarding",
		Short: "Gets email forwarding settings for the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsDNSGetEmailForwarding(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsDNSSetEmailForwarding() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSSetEmailForwardingOptions

	cmd := &cobra.Command{
		Use:   "setemailforwarding",
		Short: "Sets email forwarding for a domain name",
		RunE: func(cmd *cobra.Command, args []string) error {
			r, err := apiClient.DomainsDNSSetEmailForwarding(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.MailBox, "mailbox", nil, "Mailbox for which to set email forwarding")
	err = cmd.MarkFlagRequired("mailbox")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.ForwardTo, "forwardto", nil, "Email address to forward to")
	err = cmd.MarkFlagRequired("forwardto")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func newCmdDomainsDNSSetHosts() (*cobra.Command, error) {
	var err error
	var opts api.DomainsDNSSetHostsOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "sethosts",
		Short: "Sets DNS host records settings for the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsDNSSetHosts(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domain name")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.HostName, "hostname", nil, "Sub-domain/hostname to create the record for")
	err = cmd.MarkFlagRequired("hostname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.RecordType, "recordtype", nil, "Possible values: A, AAAA, ALIAS, CAA, CNAME, MX, MXE, NS, TXT, URL, URL301, FRAME")
	err = cmd.MarkFlagRequired("recordtype")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.Address, "address", nil, "Possible values are URL or IP address. The value for this parameter is based on RecordType.")
	err = cmd.MarkFlagRequired("address")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.TTL, "ttl", nil, "Time to live for all record types.Possible values: any value between 60 to 60000")
	err = cmd.MarkFlagRequired("ttl")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringArrayVar(&opts.MXPref, "mxpref", nil, "MX preference for host. Applicable for MX records only.")
	cmd.Flags().StringVar(&opts.Flag, "flag", "", "Is an unsigned integer between 0 and 255.")
	cmd.Flags().StringVar(&opts.Tag, "tag", "", "A non-zero sequence of US-ASCII letters and numbers in lower case.")

	return cmd, nil
}
