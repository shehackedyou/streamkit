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
		Name:        "streamkit-cli",
		Description: "streamkit command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				toolkit.X11.ActiveWindowMonitor()
				//toolkit.HandleWindowEvents()
				// aDD all the listening and event driven stuff
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			//OnExit: func(c *cli.Context) error {
			//	c.CLI.Log("on exit action")
			//	return nil
			//},
		},
	})

	fmt.Printf("toolkit:%v\n\n", toolkit)

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
