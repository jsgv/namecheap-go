package cmd

import (
	"errors"
	"strings"

	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdDomainsNS() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ns",
		Short: "Manage domains ns",
	}

	addCommand(cmd, newCmdDomainsNSCreate)
	addCommand(cmd, newCmdDomainsNSDelete)
	addCommand(cmd, newCmdDomainsNSGetInfo)
	addCommand(cmd, newCmdDomainsNSUpdate)

	return cmd, nil
}

func newCmdDomainsNSCreate() (*cobra.Command, error) {
	var err error
	var opts api.DomainsNSCreateOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new nameserver",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsNSCreate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domainname")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Nameserver, "nameserver", "", "Nameserver to create")
	err = cmd.MarkFlagRequired("nameserver")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.IP, "ip", "", "Nameserver IP address")
	err = cmd.MarkFlagRequired("ip")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
func newCmdDomainsNSDelete() (*cobra.Command, error) {
	var err error
	var opts api.DomainsNSDeleteOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a nameserver associated with the requested domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsNSDelete(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domainname")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Nameserver, "nameserver", "", "Nameserver to delete")
	err = cmd.MarkFlagRequired("nameserver")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
func newCmdDomainsNSGetInfo() (*cobra.Command, error) {
	var err error
	var opts api.DomainsNSGetInfoOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "getinfo",
		Short: "Retrieves information about a registered nameserver",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsNSGetInfo(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domainname")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Nameserver, "nameserver", "", "Nameserver")
	err = cmd.MarkFlagRequired("nameserver")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
func newCmdDomainsNSUpdate() (*cobra.Command, error) {
	var err error
	var opts api.DomainsNSUpdateOptions
	var domainname string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Updates the IP address of a registered nameserver",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			if len(domainParts) != 2 {
				return errors.New("Invalid `domainname`. Must be a valid domain name.")
			}

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsNSUpdate(opts)
			if err != nil {
				return err
			}

			return printResults(r)
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domainname")
	err = cmd.MarkFlagRequired("domainname")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.Nameserver, "nameserver", "", "Nameserver")
	err = cmd.MarkFlagRequired("nameserver")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.OldIP, "oldip", "", "Existing IP address")
	err = cmd.MarkFlagRequired("oldip")
	if err != nil {
		return nil, err
	}

	cmd.Flags().StringVar(&opts.IP, "ip", "", "New IP address")
	err = cmd.MarkFlagRequired("ip")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
