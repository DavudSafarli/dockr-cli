package cmd

import (
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
