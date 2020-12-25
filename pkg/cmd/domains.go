package cmd

import (
	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
)

func NewCmdDomains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domains",
		Short: "Manage domains",
	}

	addCommand(cmd, NewCmdDomainsDns())
	addCommand(cmd, NewCmdDomainsGetList())
	addCommand(cmd, NewCmdDomainsGetInfo())

	return cmd
}

func NewCmdDomainsGetList() *cobra.Command {
	opts := api.DomainsGetListOptions{}

	cmd := &cobra.Command{
		Use:   "getlist",
		Short: "Returns a list of domains for the particular user",
		Run: func(cmd *cobra.Command, arg []string) {
			r := apiClient.DomainsGetList(opts)
			printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.ListType, "listtype", "ALL", "Possible values are ALL, EXPIRING, or EXPIRED")
	cmd.Flags().StringVar(&opts.SearchTerm, "searchterm", "", "Keyword to look for in the domain list")
	cmd.Flags().StringVar(&opts.Page, "page", "1", "Page to return")
	cmd.Flags().StringVar(&opts.PageSize, "pagesize", "20", "Number of domains to be listed on a page. Minimum value is 10, and maximum value is 100")
	cmd.Flags().StringVar(&opts.SortBy, "sortby", "", "Possible values are NAME, NAME_DESC, EXPIREDATE, EXPIREDATE_DESC, CREATEDATE, CREATEDATE_DESC")

	return cmd
}

func NewCmdDomainsGetInfo() *cobra.Command {
	opts := api.DomainsGetInfoOptions{}

	cmd := &cobra.Command{
		Use:   "getinfo",
		Short: "Returns information about the requested domain",
		Run: func(cmd *cobra.Command, args []string) {
			r := apiClient.DomainsGetInfo(opts)
			printResults(r)
		},
	}

	cmd.Flags().StringVar(&opts.DomainName, "domainname", "", "Domain name for which domain information needs to be requested")
	cmd.Flags().StringVar(&opts.HostName, "hostname", "", "Hosted domain name for which domain information needs to be requested")

	cmd.MarkFlagRequired("domainname")

	return cmd
}
