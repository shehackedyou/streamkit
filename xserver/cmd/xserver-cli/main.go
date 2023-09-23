package main

import (
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
			//OnStart: func(c *cli.Context) error {
			//	c.CLI.Log("onStart action")
			//	return nil
			//},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			OnExit: func(c *cli.Context) error {
				c.CLI.Log("onExit action")

				activeWindow := x11.ActiveWindow()
				if xserver.IsWindowUndefined(activeWindow) {
					c.CLI.Log("returned window is undefined...\n")
				} else {
					c.CLI.Log("active window?", activeWindow.Title)
				}

				return nil
			},
		},
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
