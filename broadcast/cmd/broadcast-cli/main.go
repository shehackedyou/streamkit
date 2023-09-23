package main

import (
	"fmt"
	"os"

	broadcast "github.com/shehackedyou/streamkit/broadcast"

	cli "github.com/multiverse-os/cli"
)

func main() {
	client := &broadcast.OBS{
		Client: broadcast.Connect(broadcast.DefaultConfig()["host"]),
	}

	//show := &broadcast.Show{
	//	Scenes: make([]*show.Scene, 0),
	//}

	cmd, initErrors := cli.New(cli.App{
		Name:        "broadcast-cli",
		Description: "broadcast command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				c.CLI.Log("onStart action")

				c.CLI.Log("broadcast.OBS ", fmt.Sprintf("%v", client))
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			OnExit: func(c *cli.Context) error {
				c.CLI.Log("onExit action")

				return nil
			},
		},
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
