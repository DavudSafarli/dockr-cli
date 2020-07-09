package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
)

// ListContainers returns all containers as array in json
func ListContainers() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "returns all containers as array in json",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := GetDockerClient()

			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
				Quiet:  false,
				Size:   false,
				All:    true,
				Latest: false,
				Since:  "",
				Before: "",
				Limit:  0,
			})
			if err != nil {
				return err
			}
			var response = make(map[string]interface{}, len(containers))
			for _, container := range containers {
				id := container.ID
				response[id] = container
			}
			print(encode(response))
			return nil
		},
	}
}
