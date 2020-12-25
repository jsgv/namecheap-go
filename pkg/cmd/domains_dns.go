package cmd

import (
	"strings"

	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func NewCmdDomainsDns() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Manage domains dns",
	}

	addCommand(cmd, NewCmdDomainsDnsSetCustom())

	return cmd
}

func NewCmdDomainsDnsSetCustom() *cobra.Command {
	opts := api.DomainsDnsSetCustomOptions{}

	cmd := &cobra.Command{
		Use: "setcustom",
		Run: func(cmd *cobra.Command, args []string) {
			domainname := cmd.Flag("domainname").Value.String()
			domainParts := strings.Split(domainname, ".")

			opts.SLD = domainParts[0]
			opts.TLD = domainParts[1]

			apiClient.DomainsDnsSetCustom(opts)
		},
	}

	cmd.Flags().String("domainname", "", "Domain name")
	cmd.Flags().StringVar(&opts.Nameservers, "nameservers", "", "A comma-separated list of name servers to be associated with this domain")

	cmd.MarkFlagRequired("domainname")
	cmd.MarkFlagRequired("nameservers")

	return cmd
}
