package main

import (
	"fmt"
	"os"

	broadcast "github.com/shehackedyou/streamkit/broadcast"

	cli "github.com/multiverse-os/cli"
)

func main() {
	obs := &broadcast.OBS{
		Broadcast: broadcast.Connect(broadcast.DefaultConfig()["host"]),
		Show:      broadcast.OpenShow("she hacked you"),
	}

	cmd, initErrors := cli.New(cli.App{
		Name:        "broadcast-cli",
		Description: "broadcast command-line interface for interacting with toolkit",
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 0},
		Commands: cli.Commands(
			cli.Command{
				Name:        "list",
				Alias:       "l",
				Description: "list all obs scenes",
				Action: func(c *cli.Context) error {
					obs.ListScenes()
					return nil
				},
			},
		),
		Actions: cli.Actions{
			OnStart: func(c *cli.Context) error {
				c.CLI.Log("onStart action")

				c.CLI.Log("broadcast ", fmt.Sprintf("%v", obs))
				return nil
			},
			//Fallback: func(c *cli.Context) error {
			//	c.CLI.Log("Fallback action")
			//	return nil
			//},
			OnExit: func(c *cli.Context) error {
				c.CLI.Log("onExit action")

				fmt.Printf("scenes parsed?(%v)\n", len(obs.Show.Scenes))
				fmt.Printf("now lets iterate over OUR type of scene...\n")
				for _, s := range obs.Show.Scenes {
					fmt.Printf("Index:(%2d ) Name:( %s )\n", s.Index, s.Name)
				}

				return nil
			},
		},
	})

	if len(initErrors) == 0 {
		cmd.Parse(os.Args).Execute()
	}
}
