package cmd

import "testing"

func TestRootCmd(t *testing.T) {
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		t.Error(err)
	}
}
