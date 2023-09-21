package main

import (
	"fmt"
	"os"

	xserver "github.com/shehackedyou/streamkit/xserver"

	cli "github.com/multiverse-os/cli"
)

func main() {
	x11 := xserver.X11{
		Client: xserver.Connect(xserver.DefaultConfig()["host"]),
	}

	cmd, initErrors := cli.New(cli.App{
		Name:        "xserver-cli",
		Description: "xserver command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {

				fmt.Printf("x11.Client: %v\n", x11.Client)
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

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
