package main

import (
	"fmt"
	"os"

	streamkit "github.com/shehackedyou/streamkit"

	cli "github.com/multiverse-os/cli"
)

func main() {
	toolkit := streamkit.New()

	cmd, initErrors := cli.New(cli.App{
		Name:        "streamkit-service",
		Description: "A long-running streamkit service",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				c.CLI.Log("[onStart] performing action...")
				toolkit.X11.ActiveWindowMonitor()
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("[fallback] performing action...")
			//	return nil
			//},
			//OnExit: func(c *cli.Context) error {
			//	c.CLI.Log("[onExit] performing action...")
			//	return nil
			//},
		},
	})

	fmt.Printf("toolkit: %v\n", toolkit)

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	} else {
		fmt.Printf("errors initializing cli.App\n")
		for _, err := range initErrors {
			fmt.Printf("error(%v)\n", err)
		}
	}
}
