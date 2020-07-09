package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// Execute executes the root command.
func Execute() error {

	rootCmd := &cobra.Command{
		Use:   "dockr",
		Short: "docker api client",
		Long:  ``,
	}
	rootCmd.AddCommand(ListContainers())
	rootCmd.AddCommand(ListenEvents())

	return rootCmd.Execute()
}

// GetDockerClient returns docker client for making requests
func GetDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	check(err)
	return cli
}

func encode(v interface{}) string {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(v)
	return buf.String()
}

func print(v interface{}) {
	fmt.Fprint(os.Stdout, v)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
