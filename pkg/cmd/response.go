package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

func printResults(v interface{}) {
	// b, err := json.Marshal(v)
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(
		rootCmd.OutOrStdout(),
		string(pretty.Color(b, nil)),
	)
}
