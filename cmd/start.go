package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
)

// StartContainer starts container
func StartContainer() *cobra.Command {
	return &cobra.Command{
		Use:   "start [container id]",
		Args: cobra.ExactArgs(1),
		Short: "start a container",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := GetDockerClient()

			err := cli.ContainerStart(context.Background(), args[0], types.ContainerStartOptions{})
			if err != nil {
				return err
			}
			return nil
		},
	}
}
