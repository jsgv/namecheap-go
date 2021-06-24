package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdVersion() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display the version",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintln(cmd.OutOrStdout(), version)
			return err
		},
	}

	return cmd, nil
}
