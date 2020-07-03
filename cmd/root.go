package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute executes the root command.
func Execute() error {

	rootCmd := &cobra.Command{
		Use:   "dockr",
		Short: "docker api client",
		Long:  ``,
	}

	return rootCmd.Execute()
}

func encode(v interface{}) string {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(v)
	return buf.String()
}

func print(v interface{}) {
	fmt.Fprint(os.Stdout, encode(v))
}
