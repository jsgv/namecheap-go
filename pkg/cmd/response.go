package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

func printResults(v interface{}) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pretty.Color(b, nil)))
}
