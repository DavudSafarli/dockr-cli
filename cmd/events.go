package cmd

import (
	"context"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
)

// ListenEvents listens to events
func ListenEvents() *cobra.Command {
	return &cobra.Command{
		Use:   "listen",
		Short: "long-running process to listen docker events",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli := GetDockerClient()
			filters := filters.NewArgs()
			filters.Add("type", events.ContainerEventType)

			eventChan, _ := cli.Events(context.Background(), types.EventsOptions{
				Since:   "",
				Until:   "",
				Filters: filters,
			})
			for {
				select {
				case event := <-eventChan:
					//time.Sleep(time.Millisecond * 100)
					print(encode(event))
				}
			}
		},
	}
}
