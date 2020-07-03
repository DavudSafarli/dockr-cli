package main

import (
	"os"

	"dockr.com/go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
