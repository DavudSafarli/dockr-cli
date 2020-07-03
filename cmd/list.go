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
			var response = make([]map[string]interface{}, len(containers))
			for i, con := range containers {
				response[i] = make(map[string]interface{})
				response[i]["Command"] = con.Command
				response[i]["Created"] = con.Created
				response[i]["Id"] = con.ID
				response[i]["Image"] = con.Image
				response[i]["Names"] = con.Names
				response[i]["Ports"] = con.Ports
				response[i]["State"] = con.State
				response[i]["Status"] = con.Status
			}
			print(encode(response))
			return nil
		},
	}
}
