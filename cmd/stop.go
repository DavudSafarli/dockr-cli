package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// StopContainer stops container
func StopContainer() *cobra.Command {
	return &cobra.Command{
		Use:   "stop [container id]",
		Args: cobra.ExactArgs(1),
		Short: "stop a container",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := GetDockerClient()

			err := cli.ContainerStop(context.Background(), args[0], nil)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
