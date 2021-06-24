package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

// printResults prints with or without color depending on if the response is being
// piped.
func printResults(v interface{}) error {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		b, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(
			rootCmd.OutOrStdout(),
			string(pretty.Color(b, nil)),
		)

		if err != nil {
			return err
		}
	} else {
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(
			rootCmd.OutOrStdout(),
			string(b),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
