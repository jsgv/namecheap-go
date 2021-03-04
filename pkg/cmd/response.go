package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

func printResults(v interface{}) {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		b, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(
			rootCmd.OutOrStdout(),
			string(pretty.Color(b, nil)),
		)
	} else {
		b, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(
			rootCmd.OutOrStdout(),
			string(b),
		)
	}
}
