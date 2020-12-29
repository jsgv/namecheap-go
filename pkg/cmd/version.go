package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display the version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), version)
		},
	}

	return cmd
}
