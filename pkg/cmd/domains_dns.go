package cmd

import (
	"strings"

	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func newCmdDomainsDns() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Manage domains dns",
	}

	addCommand(cmd, newCmdDomainsDnsSetCustom())

	return cmd
}

func newCmdDomainsDnsSetCustom() *cobra.Command {
	opts := api.DomainsDnsSetCustomOptions{}
	var domainname string

	cmd := &cobra.Command{
		Use:   "setcustom",
		Short: "Sets domain to use custom DNS servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			domainParts := strings.Split(domainname, ".")

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			r, err := apiClient.DomainsDnsSetCustom(opts)
			if err != nil {
				return err
			}

			printResults(r)

			return nil
		},
	}

	cmd.Flags().StringVar(&domainname, "domainname", "", "Domain name")
	cmd.MarkFlagRequired("domainname")

	cmd.Flags().StringVar(&opts.Nameservers, "nameservers", "", "A comma-separated list of name servers to be associated with this domain")
	cmd.MarkFlagRequired("nameservers")

	return cmd
}
