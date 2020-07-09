package cmd

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
)

// ListenEvents listens to events
func ListenEvents() *cobra.Command {
	return &cobra.Command{
		Use:   "listen",
		Short: "returns all containers as array in json",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := GetDockerClient()
			eventChan, _ := cli.Events(context.Background(), types.EventsOptions{
				Since:   "",
				Until:   "",
				Filters: filters.Args{},
			})
			for {
				select {
				case event := <-eventChan:
					print(encode(event))
				}
			}
		},
	}
}
